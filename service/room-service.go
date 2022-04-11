package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type RoomService struct {
	roomRepository    *repo.RoomRepository
	floorRepostiory   *repo.FloorRepository
	requestRepository *repo.RequestRepository
}

func NewRoomService(roomRepository *repo.RoomRepository,
	floorRepostiory *repo.FloorRepository,
	requestRepository *repo.RequestRepository) *RoomService {
	return &RoomService{roomRepository: roomRepository, floorRepostiory: floorRepostiory, requestRepository: requestRepository}
}

func (s *RoomService) Read(room *model.Room) ([]model.Room, error) {
	return s.roomRepository.Read(room)
}

func (s *RoomService) Create(room *model.Room) (*model.Room, error) {
	return s.roomRepository.Create(room)
}

func (s *RoomService) Update(room *model.Room) (*model.Room, error) {
	return s.roomRepository.Update(room)
}

func (s *RoomService) Book(request *model.Request) (*model.Request, error) {
	return s.requestRepository.Create(request)
}
