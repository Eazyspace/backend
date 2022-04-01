package model

import (
	"time"
)

type Request struct {
	RequestID      int64     `json:"requestId" gorm:"primaryKey"`
	UserID         int64     `json:"userId"`
	User           User      `json:"-" gorm:"foreignKey:UserID"`
	RoomID         int64     `json:"roomId"`
	Room           Room      `json:"-"`
	StartTime      time.Time `json:"startTime,omitempty"`
	Endtime        time.Time `json:"endTime,omitempty"`
	NumberOfPeople int64     `json:"numberOfPeople,omitempty"`
	Description    string    `json:"description"`
	Status         int64     `json:"status"`
	ResponseNote   string    `json:"responseNote,omitempty"`
}
