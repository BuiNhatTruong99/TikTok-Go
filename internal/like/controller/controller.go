package controller

import (
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type likeController struct {
	likeUC like.Usecase
}

func NewLikeController(likeUC like.Usecase) *likeController {
	return &likeController{likeUC: likeUC}
}

func (lc *likeController) LikePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var likeReq entity.LikeRequest

		if err := ctx.ShouldBindJSON(&likeReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		likeRes, err := lc.likeUC.LikePost(ctx, &likeReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, likeRes)
	}
}

func (lc *likeController) UndoLikePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var likeReq entity.LikeDeleteRequest

		if err := ctx.ShouldBindJSON(&likeReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		err := lc.likeUC.UndoLikePost(ctx, &likeReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}
