package newsHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Show news
// @Schemes
// @Description Show a news
// @Tags News
// @Accept json
// @Produce json
// @Param id query string true "News identification"
// @Success 200 {object} GetNewsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /news [get]
func GetNews(ctx *gin.Context) {
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
	sendSucces(ctx, "get-news", news)
}