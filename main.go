package main

import (
	"better_calendar_backend/controllers"
	"better_calendar_backend/db"
	"better_calendar_backend/middleware"
	"better_calendar_backend/utils"
	"bufio"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var distFolder embed.FS

func main() {
	router := gin.Default()
	db.GetDB()
	router.Use(utils.CORSMiddleware())

	assets, err := fs.Sub(distFolder, "dist/assets")
	utils.ErrorHandler(err)
	router.StaticFS("assets", http.FS(assets))

	apiRouter := router.Group("/api")
	apiRouter.GET("/auth", controllers.GoogleAuth)
	template := apiRouter.Group("/template", middleware.AuthMiddleware)
	{
		template.POST("/new", controllers.CreateNewTemplate)
		template.GET("/all", controllers.GetAllTemplates)
		template.POST("/day", controllers.AddTemplateToDay)
	}

	router.GET("/", func(ctx *gin.Context) {
		html, err := distFolder.Open("dist/index.html")
		utils.ErrorHandler(err)
		defer html.Close()
		rd := bufio.NewReader(html)
		rd.WriteTo(ctx.Writer)
	})

	router.Run("localhost:10000")
}
