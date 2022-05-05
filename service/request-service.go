package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
	"github.com/mitchellh/mapstructure"
)

type RequestService struct {
	roomRepository    *repo.RoomRepository
	requestRepository *repo.RequestRepository
}

func NewRequestService(roomRepository *repo.RoomRepository,
	requestRepository *repo.RequestRepository) *RequestService {
	return &RequestService{roomRepository: roomRepository, requestRepository: requestRepository}
}

func (s *RequestService) Read(input *map[string]interface{}) ([]model.Request, error) {
	var request model.Request

	mapstructure.Decode(input, &request)
	if val, ok := (*input)["floorId"]; ok {
		if floorId, check := val.(float64); check {
			return s.requestRepository.ReadWithFloorId(int(floorId), &request)
		}
	}
	return s.requestRepository.Read(&request)
}

func (s *RequestService) Create(request *model.Request) (*model.Request, error) {
	return s.requestRepository.Create(request)
}
