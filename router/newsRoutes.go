package router

import (
	"github.com/gin-gonic/gin"
	newsHandler "main.go/handler/newsHandlers"
)

func newsRoutes(router *gin.Engine, basePath string) {
	newsHandler.InitializeNewsHandler()
	v1 := router.Group(basePath)
	{
		v1.GET("/news", newsHandler.GetNews)
		v1.GET("/listnews", newsHandler.ListNews)
		v1.POST("/news", newsHandler.CreateNews)
		v1.DELETE("/news", newsHandler.DeleteNews)
		v1.PATCH("/news", newsHandler.UpdateNews)
	}
}