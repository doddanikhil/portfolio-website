package main

import (
	"portfolio-backend/handlers"
	"portfolio-backend/middleware"
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis
	utils.InitializeRedis()

	// Create Gin Router
	r := gin.Default()

	// Middleware
	r.Use(middleware.IPLogger())

	// Routes
	r.POST("/login", handlers.LoginHandler)
	r.GET("/projects", handlers.GetProjects)
	r.POST("/projects", middleware.AuthMiddleware(), handlers.CreateProject)
	r.PUT("/projects/:id", middleware.AuthMiddleware(), handlers.UpdateProject)
	r.DELETE("/projects/:id", middleware.AuthMiddleware(), handlers.DeleteProject)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Portfolio API"})
	})
	// Start Server
	r.Run(":8080")
}
