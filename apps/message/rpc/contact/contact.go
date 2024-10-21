// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package contact

import (
	"context"

	service2 "Zhihu/application/message/rpc/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateMessageRequest     = service2.CreateMessageRequest
	Empty                    = service2.Empty
	GetFriendsListRequest    = service2.GetFriendsListRequest
	GetFriendsListResponse   = service2.GetFriendsListResponse
	GetLatestMessageRequest  = service2.GetLatestMessageRequest
	GetLatestMessageResponse = service2.GetLatestMessageResponse
	GetMessageListRequest    = service2.GetMessageListRequest
	GetMessageListResponse   = service2.GetMessageListResponse
	LoseFriendsRequest       = service2.LoseFriendsRequest
	MakeFriendsRequest       = service2.MakeFriendsRequest
	Message                  = service2.Message

	Contact interface {
		CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error)
		GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error)
		GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error)
		MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
		LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
		GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error)
	}

	defaultContact struct {
		cli zrpc.Client
	}
)

func NewContact(cli zrpc.Client) Contact {
	return &defaultContact{
		cli: cli,
	}
}

func (m *defaultContact) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.CreateMessage(ctx, in, opts...)
}

func (m *defaultContact) GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.GetLatestMessage(ctx, in, opts...)
}

func (m *defaultContact) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.GetMessageList(ctx, in, opts...)
}

func (m *defaultContact) MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.MakeFriends(ctx, in, opts...)
}

func (m *defaultContact) LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.LoseFriends(ctx, in, opts...)
}

func (m *defaultContact) GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error) {
	client := service2.NewContactClient(m.cli.Conn())
	return client.GetFriendsList(ctx, in, opts...)
}
