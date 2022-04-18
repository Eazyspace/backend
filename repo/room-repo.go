package repo

import (
	"errors"

	"github.com/Eazyspace/model"
	"gorm.io/gorm"
)

type IRoomRepository interface {
	Create(lastname, firstname string) (*model.Room, error)
	ReadByID(id int) (*model.Room, error)
	ReadAll() (*model.Room, error)
	Update(id int, lastname, firstname string) (*model.Room, error)
	Delete(id int) error
}

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepo(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}
}

func (repo *RoomRepository) Read(room *model.Room) ([]model.Room, error) {
	var foundRoom []model.Room

	// select * from Room where Room.room_code = room.room_code
	result := repo.DB.Where(room).Find(&foundRoom)

	if result.Error != nil {
		return foundRoom, result.Error
	}
	return foundRoom, nil
}

func (repo *RoomRepository) Create(room *model.Room) (*model.Room, error) {
	// check room info
	if room.MaxCapacity <= 0 ||
		room.RoomWidth <= 0 ||
		room.RoomLength <= 0 {
		return nil, errors.New("Invalid atribute (roomWidth, roomLength, maxCapacity")
	}

	// check floorId
	var foundFloor []model.Floor
	resultFloor := repo.DB.Limit(1).Where(&model.Floor{FloorID: room.FloorID}).Find(&foundFloor)
	if resultFloor.Error != nil {
		return nil, resultFloor.Error
	}
	if len(foundFloor) == 0 {
		return nil, errors.New("floorId not found")
	}

	// set room status available state
	if room.Status == 0 {
		room.Status = 1
	}

	result := repo.DB.Create(&room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

func (repo *RoomRepository) Update(room *model.Room) (*model.Room, error) {
	// check roomId
	var foundRoom []model.Room
	resultRoom := repo.DB.Limit(1).Where(&model.Room{RoomID: room.RoomID}).Find(&foundRoom)
	if resultRoom.Error != nil {
		return nil, resultRoom.Error
	}
	if len(foundRoom) == 0 {
		return nil, errors.New("roomId not found")
	}

	// set changes
	if room.FloorID != 0 {
		var foundFloor []model.Floor
		resultFloor := repo.DB.Limit(1).Where(&model.Floor{FloorID: room.FloorID}).Find(&foundFloor)
		if resultFloor.Error != nil {
			return nil, resultFloor.Error
		}
		if len(foundFloor) == 0 {
			return nil, errors.New("floorId not found")
		}
		foundRoom[0].FloorID = room.FloorID
	}
	if room.RoomName != "" {
		foundRoom[0].RoomName = room.RoomName
	}
	if room.Description != "" {
		foundRoom[0].Description = room.Description
	}
	if room.RoomWidth > 0 {
		foundRoom[0].RoomWidth = room.RoomWidth
	} else if room.RoomWidth < 0 {
		return nil, errors.New("Invalid roomWidth")
	}
	if room.RoomLength > 0 {
		foundRoom[0].RoomLength = room.RoomLength
	} else if room.RoomLength < 0 {
		return nil, errors.New("Invalid roomLength")
	}
	if room.MaxCapacity > 0 {
		foundRoom[0].MaxCapacity = room.MaxCapacity
	} else if room.MaxCapacity < 0 {
		return nil, errors.New("Invalid maxCapacity")
	}
	if 1 <= room.Status && room.Status <= 3 {
		foundRoom[0].Status = room.Status
	} else if room.Status < 0 {
		return nil, errors.New("Invalid status")
	}

	result := repo.DB.Save(&foundRoom[0])
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundRoom[0], nil
}
