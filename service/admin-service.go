package service

import (
	"errors"

	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	requestRepository *repo.RequestRepository
	userRepository    *repo.UserRepository
}

func NewAdminService(requestRepository *repo.RequestRepository, userRepository *repo.UserRepository) *AdminService {
	return &AdminService{requestRepository: requestRepository, userRepository: userRepository}
}

func (s *AdminService) UpdateStatus(request *model.Request) (*model.Request, error) {
	return s.requestRepository.UpdateStatus(request)
}

func (s *AdminService) Create(user *model.User) (*model.User, error) {
	if len(user.AcademicID) == 0 || len(user.Password) == 0 || user.Role < 1 || user.Role > 3 {
		return nil, errors.New("Missing academicId or password or invalid role")
	}

	if user.OrganizationID == 0 {
		user.OrganizationID = 1
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return s.userRepository.Create(user)
}
