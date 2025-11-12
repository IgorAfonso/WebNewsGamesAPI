package articleshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Update Article
// @Schemes
// @Description Update a Article
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "Article identification"
// @Param request body UpdateArticleRequest true "article data"
// @Success 200 {object} UpdateArticleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /article [patch]
func UpdateArticle(ctx *gin.Context) {
	request := UpdateArticleRequest{}

	ctx.BindJSON(&request)

	id := ctx.Query("id")
	if id == ""{
		sendErr(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	article := schema.Articles{}

	if err := db.Where("id = ?", id).First(&article).Error;err != nil {
		sendErr(ctx, http.StatusNotFound, "article not found")
		return
	}

	if(request.Title) != "" {article.Title = request.Title}
	if(request.Content) != "" {article.Content = request.Content}

	if err := db.Where("id = ?", id).Save(&article).Error; err != nil {
		logger.Errorf("error update article: %v", err.Error())
		sendErr(ctx, http.StatusInternalServerError, "error updating article")
		return
	}
	sendSucces(ctx, "update-news", article)
}