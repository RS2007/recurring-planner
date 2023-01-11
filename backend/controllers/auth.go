package controllers

import (
	// "better_calendar_backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func Register(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"pasword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "malformed input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"pasword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "malformed input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}
