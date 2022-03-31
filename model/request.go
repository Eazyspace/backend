package model

import (
	"time"
)

type Request struct {
	RequestID      int64     `json:"requestId" gorm:"primaryKey"`
	UserID         int64     `json:"userID"`
	User           User      `json:"-" gorm:"foreignKey:UserID"`
	RoomID         int64     `json:"roomID"`
	Room           Room      `json:"-"`
	StartTime      time.Time `json:"startTime,omitempty"`
	Endtime        time.Time `json:"endTime,omitempty"`
	NumberOfPeople int64     `json:"numberOfPeople,omitempty"`
	Description    string    `json:"description,omitempty"`
	Status         int64     `json:"status"`
	ResponseNote   string    `json:"responseNote"`
}
