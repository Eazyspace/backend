package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type FloorService struct {
	floorRepository *repo.FloorRepository
}

func NewFloorService(floorRepository *repo.FloorRepository) *FloorService {
	return &FloorService{floorRepository: floorRepository}
}

func (s *FloorService) Read(floor *model.Floor) ([]model.Floor, error) {
	return s.floorRepository.Read(floor)
}

func (s *FloorService) Create(room *model.Floor) (*model.Floor, error) {
	return s.floorRepository.Create(room)
}
