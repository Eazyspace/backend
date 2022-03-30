package model

import (
	"gorm.io/gorm"
)

type Floor struct {
	gorm.Model
	FloorCode   string `json:"floorCode"`
	FloorName   string `json:"floorName,omitempty"`
	Description string `json:"description,omitempty"`
}
