package model

type Room struct {
	RoomID      int64  `json:"roomId" gorm:"primaryKey"`
	FloorId     int64  `json:"floorId"`
	Floor       Floor  `json:"-"`
	RoomName    string `json:"roomName,omitempty"`
	RoomLength  int64  `json:"roomLength,omitempty"`
	RoomWidth   int64  `json:"roomWidth,omitempty"`
	MaxCapacity int64  `json:"maxCapacity,omitempty"`
	Description string `json:"description,omitempty"`
	Status      int64  `json:"status"`
}
