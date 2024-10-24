// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: video.proto

package videoclient

import (
	"context"

	"webshop/apps/video/rpc/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                      = video.Comment
	CommentVideoRequest          = video.CommentVideoRequest
	CommentVideoResponse         = video.CommentVideoResponse
	DeleteVideoCommentRequest    = video.DeleteVideoCommentRequest
	Empty                        = video.Empty
	FavoriteVideoRequest         = video.FavoriteVideoRequest
	GetCommentInfoRequest        = video.GetCommentInfoRequest
	GetCommentInfoResponse       = video.GetCommentInfoResponse
	GetCommentListRequest        = video.GetCommentListRequest
	GetCommentListResponse       = video.GetCommentListResponse
	GetFavoriteVideoListRequest  = video.GetFavoriteVideoListRequest
	GetFavoriteVideoListResponse = video.GetFavoriteVideoListResponse
	GetVideoListByAuthorRequest  = video.GetVideoListByAuthorRequest
	GetVideoListByAuthorResponse = video.GetVideoListByAuthorResponse
	GetVideoListRequest          = video.GetVideoListRequest
	GetVideoListResponse         = video.GetVideoListResponse
	IsFavoriteVideoRequest       = video.IsFavoriteVideoRequest
	IsFavoriteVideoResponse      = video.IsFavoriteVideoResponse
	PublishVideoRequest          = video.PublishVideoRequest
	UnFavoriteVideoRequest       = video.UnFavoriteVideoRequest
	UpdateVideoRequest           = video.UpdateVideoRequest
	VideoInfo                    = video.VideoInfo

	Video interface {
		GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error)
		PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*Empty, error)
		UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*Empty, error)
		GetVideoListByAuthor(ctx context.Context, in *GetVideoListByAuthorRequest, opts ...grpc.CallOption) (*GetVideoListByAuthorResponse, error)
		FavoriteVideo(ctx context.Context, in *FavoriteVideoRequest, opts ...grpc.CallOption) (*Empty, error)
		UnFavoriteVideo(ctx context.Context, in *UnFavoriteVideoRequest, opts ...grpc.CallOption) (*Empty, error)
		GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error)
		IsFavoriteVideo(ctx context.Context, in *IsFavoriteVideoRequest, opts ...grpc.CallOption) (*IsFavoriteVideoResponse, error)
		CommentVideo(ctx context.Context, in *CommentVideoRequest, opts ...grpc.CallOption) (*CommentVideoResponse, error)
		GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error)
		DeleteVideoComment(ctx context.Context, in *DeleteVideoCommentRequest, opts ...grpc.CallOption) (*Empty, error)
		GetCommentInfo(ctx context.Context, in *GetCommentInfoRequest, opts ...grpc.CallOption) (*GetCommentInfoResponse, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoList(ctx, in, opts...)
}

func (m *defaultVideo) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideo) UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.UpdateVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoListByAuthor(ctx context.Context, in *GetVideoListByAuthorRequest, opts ...grpc.CallOption) (*GetVideoListByAuthorResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoListByAuthor(ctx, in, opts...)
}

func (m *defaultVideo) FavoriteVideo(ctx context.Context, in *FavoriteVideoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.FavoriteVideo(ctx, in, opts...)
}

func (m *defaultVideo) UnFavoriteVideo(ctx context.Context, in *UnFavoriteVideoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.UnFavoriteVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetFavoriteVideoList(ctx, in, opts...)
}

func (m *defaultVideo) IsFavoriteVideo(ctx context.Context, in *IsFavoriteVideoRequest, opts ...grpc.CallOption) (*IsFavoriteVideoResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.IsFavoriteVideo(ctx, in, opts...)
}

func (m *defaultVideo) CommentVideo(ctx context.Context, in *CommentVideoRequest, opts ...grpc.CallOption) (*CommentVideoResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.CommentVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}

func (m *defaultVideo) DeleteVideoComment(ctx context.Context, in *DeleteVideoCommentRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.DeleteVideoComment(ctx, in, opts...)
}

func (m *defaultVideo) GetCommentInfo(ctx context.Context, in *GetCommentInfoRequest, opts ...grpc.CallOption) (*GetCommentInfoResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetCommentInfo(ctx, in, opts...)
}
