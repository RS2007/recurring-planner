package main

import (
	"better_calendar_backend/controllers"
	"better_calendar_backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.GetDB()
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
	router.Run("localhost:5000")
}
