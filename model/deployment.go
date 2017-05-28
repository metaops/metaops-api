package model

import (
	"time"
)

type Deployment struct {
	ID     string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	AppID  string `gorm:"type:uuid"`
	Status string `gorm:"type:varchar(100)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
