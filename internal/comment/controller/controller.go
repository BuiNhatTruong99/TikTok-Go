package controller

import (
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type commentController struct {
	commentUC comment.Usecase
}

func NewCommentController(commentUC comment.Usecase) *commentController {
	return &commentController{commentUC: commentUC}
}

func (cmc *commentController) CommentPost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var commentReq entity.CommentReqest

		if err := ctx.ShouldBindJSON(&commentReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := cmc.commentUC.CommentPost(ctx, &commentReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}

func (cmc *commentController) DeleteComment() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var commentDeleteReq entity.CommentDeleteReqest

		if err := ctx.ShouldBindJSON(&commentDeleteReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := cmc.commentUC.DeleteComment(ctx, &commentDeleteReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}
