package articleshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary List Articles
// @Schemes
// @Description List all articles
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {object} ListArticleResponse
// @Failure 500 {object} ErrorResponse
// @Router /articles [get]
func ListArticles(ctx *gin.Context) {
	article := []schema.Articles{}

	if err := db.Find(&article).Error; err != nil{
		sendErr(ctx, http.StatusInternalServerError, "error listing articles")
		return
	}
	sendSucces(ctx, "list-articles", article)
}