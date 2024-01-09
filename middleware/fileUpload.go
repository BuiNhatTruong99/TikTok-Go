package middleware

import (
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: "couldnt get file"})
			return
		}
		defer file.Close()

		c.Set("filePath", header.Filename)
		c.Set("file", file)

		c.Next()
	}
}
