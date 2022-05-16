package service

import (
	"errors"
	"fmt"

	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
	"github.com/mitchellh/mapstructure"
)

type RequestService struct {
	userRepository    *repo.UserRepository
	roomRepository    *repo.RoomRepository
	requestRepository *repo.RequestRepository
}

func NewRequestService(roomRepository *repo.RoomRepository,
	requestRepository *repo.RequestRepository, userRepository *repo.UserRepository) *RequestService {
	return &RequestService{roomRepository: roomRepository, requestRepository: requestRepository, userRepository: userRepository}
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

func (s *RequestService) CheckIn(request *model.Request, userId int64) (bool, error) {

	if request.RequestID == 0 {
		request.RequestID = -1
	}
	req, err := s.requestRepository.Read(request)
	fmt.Println(len(req))
	if len(req) == 0 {
		return false, errors.New("not found")
	}
	if req[0].UserID != userId {
		return false, errors.New("unauthorized")
	}
	if err != nil {
		return false, err
	}

	return s.requestRepository.CheckIn(request)
}
