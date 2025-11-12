package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/docs"
)

func SwagRoute(router *gin.Engine, basePath string) {
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}