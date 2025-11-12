package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	ID         uuid.UUID
	Title      string
	Content    string
}

type ArticlesResponse struct {
	ID uint 			`json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Title string 		`json:"title"`
	Content string 		`json:"content"`
}