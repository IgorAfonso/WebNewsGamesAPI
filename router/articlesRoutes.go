package router

import (
	"github.com/gin-gonic/gin"
	articlesHandlers "main.go/handler/articlesHandlers"
)

func articlesRoutes(router *gin.Engine, basePath string) {
	articlesHandlers.InitializeNewsHandler()
	v1 := router.Group(basePath)
	{
		v1.GET("/article", articlesHandlers.GetArticle)
		v1.GET("/articles", articlesHandlers.ListArticles)
		v1.POST("/article", articlesHandlers.CreateArticle)
		v1.DELETE("/article", articlesHandlers.DeleteArticle)
		v1.PATCH("/article", articlesHandlers.UpdateArticle)
	}
}