package controller

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

type userController struct {
	userUC user.UseCase
	cfg    *config.Config
}

func NewUserController(userUC user.UseCase, cfg *config.Config) *userController {
	return &userController{userUC: userUC, cfg: cfg}
}

func (uc *userController) ChangeAvatar() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("user-id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if _, err := uc.userUC.GetUserByID(ctx, int64(userID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		file, _, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: "could not find file in request"})
			return
		}
		defer file.Close()

		var avatarReq entity.AvatarRequest

		avatarUrl, err := utils.UploadToCloudinary(file.(multipart.File), uc.cfg,
			uc.cfg.Cloudinary.CloudUploadFolderAvatar)

		avatarReq.AvatarUrl = avatarUrl

		if err := uc.userUC.ChangeAvatar(ctx, int64(userID), &avatarReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}

func (uc *userController) UpdateProfile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("user-id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		var profileReq entity.ProfileRequest

		if err := ctx.ShouldBindJSON(&profileReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := uc.userUC.UpdateProfile(ctx, int64(userID), &profileReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}
