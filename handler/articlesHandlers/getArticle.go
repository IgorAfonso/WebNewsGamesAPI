package articleshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Show Article
// @Schemes
// @Description Show a Article
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "Article identification"
// @Success 200 {object} GetArticleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /article [get]
func GetArticle(ctx *gin.Context) {
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
	sendSucces(ctx, "get-article", article)
}