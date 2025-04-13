package main

import (
	"hello/config"
	"hello/models"
	"hello/routes"
	"hello/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
