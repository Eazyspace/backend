package model

import (
	"time"
)

type Token struct {
	UserID      int64     `json:"userId" gorm:"primaryKey"`
	ExpireTime  time.Time `json:"expireTime,omitempty"`
	TokenString string    `json:"tokenString,omitempty"`
}
