package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	basePath := "/api/v1"
	router := gin.Default()

	// Initialize Swagger
	SwagRoute(router, basePath)

	// Declare News Routes
	newsRoutes(router, basePath)

	// Declare Articles Routes
	articlesRoutes(router, basePath)

	//Initialzie Aplication in default port
	router.Run(":8080")
}