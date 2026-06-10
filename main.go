package main

import (
	"URLSHORTNER/config"
	"URLSHORTNER/models"
	"URLSHORTNER/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Url{})

	r := routes.SetupRouter()

	r.Run()

}
