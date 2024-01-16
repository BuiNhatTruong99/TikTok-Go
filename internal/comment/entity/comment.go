package entity

import "time"

type Comment struct {
	ID        int64      `json:"id" gorm:"column:id;primaryKey"`
	UserID    int64      `json:"user_id" gorm:"column:user_id"`
	PostID    int64      `json:"post_id" gorm:"column:post_id"`
	Text      string     `json:"text" gorm:"column:text"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

type CommentReqest struct {
	UserID int64  `json:"user_id" gorm:"column:user_id"`
	PostID int64  `json:"post_id" gorm:"column:post_id"`
	Text   string `json:"text" gorm:"column:text"`
}

type CommentDeleteReqest struct {
	CommentID int64 `json:"id" gorm:"column:id;primaryKey"`
	UserID    int64 `json:"user_id" gorm:"column:user_id"`
}

func TableName() string {
	return "comments"
}
