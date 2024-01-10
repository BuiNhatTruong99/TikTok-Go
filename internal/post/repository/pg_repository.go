package repository

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.Repository {
	return &postRepo{db: db}
}

func (p *postRepo) CreatePost(ctx context.Context, request *entity.PostRequest) error {
	if err := p.db.Table(entity.TableName()).WithContext(ctx).Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (p *postRepo) GetPostByID(ctx context.Context, postID int64) (*entity.Post, error) {
	var getPost entity.Post

	if err := p.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", postID).First(&getPost).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &getPost, nil
}

func (p *postRepo) GetPostByUserID(ctx context.Context, userID int64) ([]entity.Post, error) {
	var posts []entity.Post
	if err := p.db.Table(entity.TableName()).WithContext(ctx).Where("user_id = ?",
		userID).Find(&posts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no post has found")
		}
		return nil, err
	}
	return posts, nil
}

func (p *postRepo) DeletePostByID(ctx context.Context, postID int64) error {
	if err := p.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?",
		postID).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (p *postRepo) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	var posts []entity.Post
	if err := p.db.Table(entity.TableName()).WithContext(ctx).Find(&posts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no post has found")
		}
		return nil, err
	}
	return posts, nil
}
