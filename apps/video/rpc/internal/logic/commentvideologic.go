package logic

import (
	"context"
	"fmt"

	"gorm.io/gorm/clause"
	"webshop/apps/video/mq"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CommentVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentVideoLogic {
	return &CommentVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const HotVideo = 1

func (l *CommentVideoLogic) CommentVideo(in *video.CommentVideoRequest) (*video.CommentVideoResponse, error) {
	//create comment record
	comment := model.Comment{
		VideoId: in.VideoId,
		UserId:  in.UserId,
		Content: in.Content,
	}

	//start the transaction
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}
		// update the number of comments
		result, err := l.svcCtx.Redis.Exists(l.ctx, genVideoInfoCacheKey(in.VideoId)).Result()
		if err != nil {
			return err
		}
		// if it is a hot video (in the cache), only update the cache and hand it over to the timed task to update the database
		if result == HotVideo {
			task, err := mq.NewAddCacheValueTask(genVideoInfoCacheKey(in.VideoId), "CommentCount", 1)
			if err != nil {
				return err
			}
			// add the task to the queue
			if _, err := l.svcCtx.AsynqClient.Enqueue(task); err != nil {
				return err
			}
		} else {
			if err != nil {
				l.Logger.Errorf("failed to check key: %v", err)
			}
			//use Clause to lock the row to prevent other transactions from modifying the row
			if err := tx.Model(&model.Video{}).
				Where("id = ?", in.VideoId).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).
				Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.Logger.Errorf("failed to create comment: %v", err)
		return nil, xerr.NewErrCode(xerr.VideoFailToCreateComment)
	}

	return &video.CommentVideoResponse{
		Id:          int64(comment.ID),
		UserId:      comment.UserId,
		Content:     comment.Content,
		CreatedTime: comment.CreatedAt.Unix(),
	}, nil
}

func genVideoInfoCacheKey(videoId int64) string {
	return fmt.Sprintf("biz#video#info#%d", videoId)
}
