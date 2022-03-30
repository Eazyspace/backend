package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserCode    string    `json:"userCode"`
	ExpireTime  time.Time `json:"expireTime,omitempty"`
	TokenString string    `json:"tokenString,omitempty"`
}
