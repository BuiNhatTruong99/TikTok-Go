package utils

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"strings"
)

func UploadToCloudinary(file multipart.File, cfg *config.Config) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary(cfg)
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		Folder: cfg.Cloudinary.CloudUploadFolder,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	videoUrl := result.SecureURL
	return videoUrl, nil
}

func RemoveFromCloudinary(videoURL string, cfg *config.Config) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary(cfg)
	if err != nil {
		return "", err
	}

	parts := strings.Split(videoURL, "/")
	publicID := strings.TrimSuffix(parts[len(parts)-1], ".mp4")

	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     cfg.Cloudinary.CloudUploadFolder + "/" + publicID,
		ResourceType: "video",
	})
	if err != nil {
		return "", err
	}

	return resp.Result, nil
}
