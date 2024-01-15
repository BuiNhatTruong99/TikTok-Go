package entity

import "time"

type Like struct {
	ID        int64      `json:"id" gorm:"column:id;primaryKey"`
	UserID    int64      `json:"user_id" gorm:"column:user_id"`
	PostID    int64      `json:"post_id" gorm:"column:post_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

type LikeRequest struct {
	UserID int64 `json:"user_id" gorm:"column:user_id"`
	PostID int64 `json:"post_id" gorm:"column:post_id"`
}

type LikeDeleteRequest struct {
	UserID int64 `json:"user_id" gorm:"column:user_id"`
	LikeID int64 `json:"id" gorm:"column:id"`
}

type LikeResponse struct {
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
}

func TableName() string {
	return "likes"
}
