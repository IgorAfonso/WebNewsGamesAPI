package newsHandler

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

type CreateNewsResponse struct {
	Message string 					`json:"message"`
	Data schema.NewsResponse 		`json:"data"`
}

type DeleteNewsResponse struct{
	Message string 					`json:"message"`
	Data schema.NewsResponse 		`json:"data"`
}

type GetNewsResponse struct{
	Message string 					`json:"message"`
	Data schema.NewsResponse 		`json:"data"`
}

type ListNewsResponse struct{
	Message string 					`json:"message"`
	Data []schema.NewsResponse 		`json:"data"`
}

type UpdateNewsResponse struct{
	Message string 					`json:"message"`
	Data schema.NewsResponse 		`json:"data"`
}