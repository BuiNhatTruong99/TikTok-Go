package usecase

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
)

type postUC struct {
	postRepository post.Repository
}

func NewPostUsecase(postRepo post.Repository) post.UseCase {
	return &postUC{postRepository: postRepo}
}

func (p *postUC) CreatePost(ctx context.Context, request *entity.PostRequest) error {
	return p.postRepository.CreatePost(ctx, request)
}

func (p *postUC) DeletePost(ctx context.Context, postID int64) error {
	return p.postRepository.DeletePostByID(ctx, postID)
}

func (p *postUC) GetPostByID(ctx context.Context, postID int64) (*entity.Post, error) {
	return p.postRepository.GetPostByID(ctx, postID)
}

func (p *postUC) GetPostsByUserID(ctx context.Context, userID int64) ([]entity.Post, error) {
	return p.postRepository.GetPostByUserID(ctx, userID)
}

func (p *postUC) GetAllPost(ctx context.Context) ([]entity.Post, error) {
	return p.postRepository.GetAllPosts(ctx)
}
