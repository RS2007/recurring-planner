package main

import (
	"better_calendar_backend/controllers"
	"better_calendar_backend/db"
	"better_calendar_backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.GetDB()
	router.Use(utils.CORSMiddleware())
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
	router.Run("localhost:5000")
}
