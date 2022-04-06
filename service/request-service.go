package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type RequestService struct {
	roomRepository    *repo.RoomRepository
	requestRepository *repo.RequestRepository
}

func NewRequestService(roomRepository *repo.RoomRepository,
	requestRepository *repo.RequestRepository) *RequestService {
	return &RequestService{roomRepository: roomRepository, requestRepository: requestRepository}
}

func (s *RequestService) Read(request *model.Request) ([]model.Request, error) {
	return s.requestRepository.Read(request)
}

func (s *RequestService) Create(request *model.Request) (*model.Request, error) {
	return s.requestRepository.Create(request)
}
