package newsHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Create News
// @Description Create a news for the app
// @Tags News
// @Accept json
// @Produce json
// @Param request body CreateNewsObject true "Request Body"
// @Success 201 {object} CreateNewsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /news [post]
func CreateNews(ctx *gin.Context) {
	request := CreateNewsObject{}

	ctx.BindJSON(&request)

	if err := createNewsValidations(request); err != nil{
		logger.Errorf("validation error: %+v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	news := schema.News{
		ID: uuid.New(),
		Title: request.Title,
		FirstContent: request.FirstContent,
		SecondContent: request.SecondContent,
		ThirdContent: request.ThirdContent,
	}

	if err := db.Create(&news).Error; err != nil {
		logger.Infof("error creating news %+v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendPostSucces(ctx, "creating-news", news)
}