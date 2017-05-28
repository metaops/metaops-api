package model

import (
	"time"
)

type Node struct {
	ID    string `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	AppID string `gorm:"type:uuid" json:"appId"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
