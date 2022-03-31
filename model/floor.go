package model

type Floor struct {
	FloorID     int64  `json:"floorId" gorm:"primaryKey"`
	FloorName   string `json:"floorName,omitempty"`
	Description string `json:"description,omitempty"`
}
