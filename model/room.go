package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomCode    string `json:"roomCode"`
	RoomName    string `json:"roomName,omitempty"`
	RoomLength  int64  `json:"roomLength,omitempty"`
	RoomWidth   int64  `json:"roomWidth,omitempty"`
	MaxCapacity int64  `json:"maxCapacity,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}
