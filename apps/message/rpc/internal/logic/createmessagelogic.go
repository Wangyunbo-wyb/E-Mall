package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"webshop/apps/message/rpc/internal/model"
	"webshop/apps/message/rpc/internal/svc"
	"webshop/apps/message/rpc/service"
	"webshop/pkg/xerr"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMessageLogic) CreateMessage(in *service.CreateMessageRequest) (*service.Empty, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//创建并增加消息记录
		message := model.Message{
			FromId:   in.FromId,
			ToUserId: in.ToId,
			Content:  in.Content,
		}

		if err := l.svcCtx.DB.Create(&message).Error; err != nil {
			l.Logger.Errorf("[CreateMessage] create message failed: %v", err)
			return xerr.NewErrCode(xerr.MessageDataBaseError)
		}

		return nil
	})
	if err != nil {
		l.Logger.Info("[CreateMessage] the transaction failed: %v", err)
		return nil, xerr.NewErrCode(xerr.MessageDataBaseError)
	}
	return &service.Empty{}, nil
}
