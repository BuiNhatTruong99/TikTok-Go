package entity

import "time"

type Post struct {
	ID        int64      `json:"id" gorm:"column:id;primaryKey"`
	UserID    int64      `json:"user_id" gorm:"column:user_id"`
	VideoUrl  string     `json:"video_url" gorm:"column:video_url"`
	Caption   string     `json:"caption" gorm:"column:caption"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

type PostRequest struct {
	UserID   int64  `json:"user_id" gorm:"column:user_id"`
	VideoUrl string `json:"video_url" gorm:"column:video_url"`
	Caption  string `json:"caption" gorm:"column:caption"`
}

func TableName() string {
	return "posts"
}
