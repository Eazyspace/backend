package repo

import (
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
	result := repo.DB.Create(&room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}
