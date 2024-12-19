package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	// Mock projects (replace with database/Redis)
	projects := []map[string]string{
		{"id": "1", "title": "Project 1", "description": "Description 1"},
		{"id": "2", "title": "Project 2", "description": "Description 2"},
	}
	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	// Mock create project
	c.JSON(http.StatusCreated, gin.H{"message": "Project created"})
}

func UpdateProject(c *gin.Context) {
	// Mock update project
	c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
}

func DeleteProject(c *gin.Context) {
	// Mock delete project
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
