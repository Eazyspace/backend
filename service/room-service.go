package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type IRoomService interface {
	Create(lastname, firstname string) (*model.Room, error)
	ReadByID(id int) (*model.Room, error)
	ReadAll() (*model.Room, error)
	Update(id int, lastname, firstname string) (*model.Room, error)
	Delete(id int) error
}

var roomService IRoomService

type RoomService struct {
	roomRepository *repo.RoomRepository
}

func NewRoomService(roomRepository *repo.RoomRepository) *RoomService {
	return &RoomService{roomRepository: roomRepository}
}
func (s *RoomService) Read(room *model.Room) ([]model.Room, error) {
	return s.roomRepository.Read(room)

}
func (s *RoomService) Create(room *model.Room) (*model.Room, error) {
	return s.roomRepository.Create(room)
}
