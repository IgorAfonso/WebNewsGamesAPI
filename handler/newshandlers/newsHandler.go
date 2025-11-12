package newsHandler

import (
	"gorm.io/gorm"
	"main.go/config"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeNewsHandler(){
	logger = config.GetLogger("news")
	db = config.GetPostgres()
}