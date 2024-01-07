package controller

import (
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type postController struct {
	postUC post.UseCase
}

func NewPostController(postUC post.UseCase) *postController {
	return &postController{postUC: postUC}
}

func (p *postController) CreatePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var postReq entity.PostRequest

		if err := ctx.ShouldBindJSON(&postReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := p.postUC.CreatePost(ctx, &postReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}

func (p *postController) DeletePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		postID, err := strconv.Atoi(ctx.Param("post-id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := p.postUC.DeletePost(ctx, int64(postID)); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}
