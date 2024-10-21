package logic

import (
	"context"

	"Zhihu/application/message/rpc/contact"
	"Zhihu/application/message/rpc/internal/code"
	"Zhihu/application/message/rpc/internal/model"
	"Zhihu/application/message/rpc/internal/svc"
	"Zhihu/application/message/rpc/service"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *service.GetMessageListRequest) (*service.GetMessageListResponse, error) {
	var messages []model.Message
	err := l.svcCtx.DB.Where("from_id = ?", in.FromId).Where("to_user_id = ?", in.ToId).Find(&messages).Error
	if err != nil {
		l.Logger.Errorf("[GetMessageList] get message list failed: %v", err)
		return nil, code.MessageDataBaseError
	}

	var messageList []*contact.Message
	for _, message := range messages {
		messageList = append(messageList, &contact.Message{
			Id:         int64(message.ID),
			Content:    message.Content,
			CreateTime: message.CreatedAt.Unix(),
			FromId:     message.FromId,
			ToId:       message.ToUserId,
		})
	}
	return &contact.GetMessageListResponse{
		Messages: messageList,
	}, nil
}
