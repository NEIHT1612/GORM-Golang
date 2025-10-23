package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/models"
	"example.com/m/v2/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.InitDB()
	models.RunMigrations()

	server := gin.Default()
	route.RegisterRoutes(server)
	server.Run(":8080")
}