package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/schema"
)

func InitializePostgres() (*gorm.DB, error){
	if err := godotenv.Load(); err != nil {
		println("Error to load .env file")
	}
	
	conectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(conectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to conect to database. %s", err)
		fmt.Printf("%s", conectionString)
		return nil, err
	}

	err = db.AutoMigrate(&schema.News{}, &schema.Articles{})
	if err != nil {
		fmt.Printf("PostGres automigration error. %s", err)
	}

	return db, nil
}