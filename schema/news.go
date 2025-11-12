package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID uuid.UUID
	Title string
	FirstContent string
	SecondContent string
	ThirdContent string
}

type NewsResponse struct {
	ID uint 				`json:"id"`
	CreatedAt time.Time 	`json:"createdAt"`
	UpdatedAt time.Time 	`json:"updatedAt"`
	DeletedAt time.Time 	`json:"deletedAt,omitempty"`
	FirstContent string 	`json:"firstContent"`
	SecondContent string 	`json:"secondContent"`
	ThirdContent string 	`json:"thirdContent"`
}