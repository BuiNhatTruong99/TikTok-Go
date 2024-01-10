package controller

import (
	"encoding/json"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

type postController struct {
	postUC post.UseCase
	cfg    *config.Config
}

func NewPostController(postUC post.UseCase, cfg *config.Config) *postController {
	return &postController{postUC: postUC, cfg: cfg}
}

func (p *postController) CreatePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		file, _, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: "could not find file in request"})
			return
		}
		defer file.Close()

		var postReq entity.PostRequest

		postString := ctx.Request.FormValue("post")
		err = json.Unmarshal([]byte(postString), &postReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}
		videoUrl, err := utils.UploadToCloudinary(file.(multipart.File), p.cfg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		postReq.VideoUrl = videoUrl

		if err := p.postUC.CreatePost(ctx, &postReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}

func (p *postController) GetAllPosts() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		posts, err := p.postUC.GetAllPost(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}

func (p *postController) GetPostsByUserID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("user-id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		posts, err := p.postUC.GetPostsByUserID(ctx, int64(userID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, posts)
	}
}

func (p *postController) DeletePost() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		postID, err := strconv.Atoi(ctx.Param("post-id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		post, err := p.postUC.GetPostByID(ctx, int64(postID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		res, err := utils.RemoveFromCloudinary(post.VideoUrl, p.cfg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := p.postUC.DeletePost(ctx, int64(postID)); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(res))
	}
}
