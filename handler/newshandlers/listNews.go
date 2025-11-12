package newsHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary List news
// @Schemes
// @Description List all news
// @Tags News
// @Accept json
// @Produce json
// @Success 200 {object} ListNewsResponse
// @Failure 500 {object} ErrorResponse
// @Router /listnews [get]
func ListNews(ctx *gin.Context) {
	news := []schema.News{}

	if err := db.Find(&news).Error; err != nil{
		sendErr(ctx, http.StatusInternalServerError, "error listing news")
		return
	}
	sendSucces(ctx, "list-news", news)
}