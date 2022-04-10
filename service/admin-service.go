package service

import (
	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
)

type AdminService struct {
	requestRepository *repo.RequestRepository
}

func NewAdminService(requestRepository *repo.RequestRepository) *AdminService {
	return &AdminService{requestRepository: requestRepository}
}

func (s *AdminService) UpdateStatus(request *model.Request) (*model.Request, error) {
	return s.requestRepository.UpdateStatus(request)
}
