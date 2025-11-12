package articleshandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

func sendErr(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendPostSucces(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation %s successful", op),
		"data":    data,
	})
}

func sendSucces(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateArticleResponse struct {
	Message string 						`json:"message"`
	Data schema.ArticlesResponse 		`json:"data"`
}

type DeleteArticleResponse struct{
	Message string 						`json:"message"`
	Data schema.ArticlesResponse 		`json:"data"`
}

type GetArticleResponse struct{
	Message string 						`json:"message"`
	Data schema.ArticlesResponse 		`json:"data"`
}

type ListArticleResponse struct{
	Message string 						`json:"message"`
	Data []schema.ArticlesResponse 		`json:"data"`
}

type UpdateArticleResponse struct{
	Message string 						`json:"message"`
	Data schema.ArticlesResponse 		`json:"data"`
}