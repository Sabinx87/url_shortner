package controllers

import (
	"URLSHORTNER/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func ShortenURL(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url is require",
		})
		return
	}
	shortend, err := services.CreateShortUrl(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "fail to short url",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"short":    shortend.Short,
		"original": shortend.Original,
	})

}

func RedirectToOriginal(c *gin.Context) {
	short := c.Param("short")
	url, err := services.GetOriginalURL(short)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "short url not found",
		})
	}
	c.Redirect(http.StatusMovedPermanently, url.Original)

}
