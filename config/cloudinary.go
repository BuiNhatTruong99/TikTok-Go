package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary(cfg *Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.CloudAPIKey, cfg.Cloudinary.CloudAPISecretKey)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
