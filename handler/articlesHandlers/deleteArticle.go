package articleshandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/schema"
)

// @BasePath /api/v1

// @Summary Delete Article
// @Schemas
// @Description Delete a Article
// @Tags Article
// @Accept json
// @Produce json
// @Param id query string true "Article identification"
// @Success 200 {object} DeleteArticleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /article [delete]
func DeleteArticle(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		sendErr(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	article := schema.Articles{}
	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("article with id: %s not found", id))
		return
	}

	if err := db.Where("id = ?", id).Delete(&article).Error; err != nil {
		sendErr(ctx, http.StatusInternalServerError, 
			fmt.Sprintf("error deleting article with id: %s", id))
			return
	}

	sendSucces(ctx, "delete-article", article)
}