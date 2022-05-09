package service

import (
	"errors"
	"time"

	"github.com/Eazyspace/model"
	"github.com/Eazyspace/repo"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repo.UserRepository
}

func NewUserService(userRepository *repo.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Read(user *model.User) ([]model.User, error) {
	return s.userRepository.Read(user)
}

func (s *UserService) Create(user *model.User) (*model.User, error) {

	if len(user.AcademicID) == 0 || len(user.Password) == 0 {
		return nil, errors.New("Missing academicId or password")
	}

	if user.OrganizationID == 0 {
		user.OrganizationID = 1
	}
	user.Role = 1
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return s.userRepository.Create(user)
}

func (s *UserService) Login(user *model.User) (*string, error) {
	// Hash password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(string(hashedPassword))
	password := user.Password
	user.Password = ""

	// find user account match username and password
	var foundUsers []model.User
	foundUsers, resultUser := s.userRepository.Read(user)
	if resultUser != nil {
		return nil, echo.ErrUnauthorized
	}
	if len(foundUsers) == 0 {
		return nil, errors.New("invalid username")
	}
	if bcrypt.CompareHashAndPassword([]byte(foundUsers[0].Password), []byte(password)) != nil {
		return nil, errors.New("invalid password")
	}

	// Set custom claims
	claims := &model.Token{
		UserID:         foundUsers[0].UserID,
		Role:           foundUsers[0].Role,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (s *UserService) SetAvatar(user *model.User) (*model.User, error) {
	return s.userRepository.SetAvatar(user)
}
