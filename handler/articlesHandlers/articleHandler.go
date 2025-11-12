package articleshandlers

import (
	"gorm.io/gorm"
	"main.go/config"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeNewsHandler(){
	logger = config.GetLogger("article")
	db = config.GetPostgres()
}