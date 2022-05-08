package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	UserID      int64     `json:"userId" gorm:"primaryKey"`
	User        User      `gorm:"foreignKey:UserID"`
	ExpireTime  time.Time `json:"expireTime,omitempty"`
	Role        int64     `josn:"role,omitempty"`
	TokenString string    `json:"tokenString,omitempty"`
	jwt.StandardClaims
}
