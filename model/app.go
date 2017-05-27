package model

import (
	"time"
)

type App struct {
	ID   string `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
