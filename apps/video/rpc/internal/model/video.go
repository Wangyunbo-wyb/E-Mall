package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId      uint64 `gorm:"not null;index"`
	Title         string `gorm:"not null;index"`
	PlayUrl       string `gorm:"not null"`
	CoverUrl      string `gorm:"not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64

	// has many
	Comments  []Comment
	Favorites []Favorite
}

const (
	PopularVideoStandard = 1000 // the video with more than 1000 likes or 1000 comments becomes a popular video and has special processing
)

func IsPopularVideo(favoriteCount, commentCount int64) bool {
	return favoriteCount >= PopularVideoStandard || commentCount >= PopularVideoStandard
}

type Comment struct {
	gorm.Model
	UserId  uint64
	VideoId int64
	Content string `gorm:"not null"`
}

type Favorite struct {
	gorm.Model
	UserId  uint64 `gorm:"column:user_id;"`
	VideoId int64  `gorm:"column:video_id;"`

	// belongs to
	Video Video
}
