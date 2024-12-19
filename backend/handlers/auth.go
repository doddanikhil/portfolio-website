package handlers

import (
	"net/http"
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// Mock user credentials (replace with database validation)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "admin" && password == "password" {
		token, _ := utils.GenerateJWT(username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
