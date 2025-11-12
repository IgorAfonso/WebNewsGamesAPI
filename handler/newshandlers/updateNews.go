package newsHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Update news
// @Schemes
// @Description Update a news
// @Tags News
// @Accept json
// @Produce json
// @Param id query string true "News identification"
// @Param request body UpdateNewsObject true "News data"
// @Success 200 {object} UpdateNewsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /news [patch]
func UpdateNews(ctx *gin.Context) {
	request := UpdateNewsObject{}

	ctx.BindJSON(&request)

	id := ctx.Query("id")
	if id == ""{
		sendErr(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	news := schema.News{}

	if err := db.Where("id = ?", id).First(&news).Error;err != nil {
		sendErr(ctx, http.StatusNotFound, "news not found")
		return
	}

	if(request.FirstContent) != "" {news.FirstContent = request.FirstContent}
	if(request.SecondContent) != "" {news.SecondContent = request.SecondContent}
	if(request.ThirdContent) != "" {news.ThirdContent = request.ThirdContent}
	if(request.Title) != "" {news.Title = request.Title}

	if err := db.Where("id = ?", id).Save(&news).Error; err != nil {
		logger.Errorf("error update news: %v", err.Error())
		sendErr(ctx, http.StatusInternalServerError, "error updating news")
		return
	}
	sendSucces(ctx, "update-news", news)
}