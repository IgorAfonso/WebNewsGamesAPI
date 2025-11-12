package articleshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Create Articles
// @Description Create a Article for the app
// @Tags Article
// @Accept json
// @Produce json
// @Param request body CreateArticleRequest true "Request Body"
// @Success 201 {object} CreateArticleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /article [post]
func CreateArticle(ctx *gin.Context) {
	request := CreateArticleRequest{}

	ctx.BindJSON(&request)

	if err := createArticleValidations(request); err != nil{
		logger.Errorf("validation error: %+v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	article := schema.Articles{
		ID: uuid.New(),
		Title: request.Title,
		Content: request.Content,
	}

	if err := db.Create(&article).Error; err != nil {
		logger.Infof("error creating article %+v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendPostSucces(ctx, "creating-article", article)
}