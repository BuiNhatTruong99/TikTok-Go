package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comment.Repository {
	return &commentRepo{db: db}
}

func (c *commentRepo) CreateComment(ctx context.Context, commentReq *entity.CommentReqest) error {
	if err := c.db.Table(entity.TableName()).WithContext(ctx).Create(&commentReq).Error; err != nil {
		return err
	}
	return nil
}

func (c *commentRepo) DeleteComment(ctx context.Context, commentID int64) error {
	if err := c.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", commentID).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (c *commentRepo) GetCommentByID(ctx context.Context, commentID int64) (*entity.Comment, error) {
	var getComment entity.Comment
	if err := c.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?",
		commentID).First(&getComment).Error; err != nil {
		return nil, err
	}
	return &getComment, nil
}
