package services

import (
	"URLSHORTNER/config"
	"URLSHORTNER/models"
	"URLSHORTNER/utils"
)

func CreateShortUrl(original string) (models.Url, error) {
	shortCode := utils.GenerateShortCode(6)

	url := models.Url{
		Original: original,
		Short:    shortCode,
	}
	result := config.DB.Create(&url)

	return url, result.Error
}
