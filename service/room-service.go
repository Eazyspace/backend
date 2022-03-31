package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type RoomService struct {
	roomRepository  *repo.RoomRepository
	floorRepostiory *repo.FloorRepository
}

func NewRoomService(roomRepository *repo.RoomRepository,
	floorRepostiory *repo.FloorRepository) *RoomService {
	return &RoomService{roomRepository: roomRepository, floorRepostiory: floorRepostiory}
}

func (s *RoomService) Read(room *model.Room) ([]model.Room, error) {
	return s.roomRepository.Read(room)
}

func (s *RoomService) Create(room *model.Room) (*model.Room, error) {
	return s.roomRepository.Create(room)
}
