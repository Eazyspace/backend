package repo

import (
	"errors"

	"github.com/Eazyspace/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(lastname, firstname string) (*model.User, error)
	ReadByID(id int) (*model.User, error)
	ReadAll() (*model.User, error)
	Update(id int, lastname, firstname string) (*model.User, error)
	Delete(id int) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) ReadAll() ([]model.User, error) {
	var foundUsers []model.User

	result := repo.DB.Find(&foundUsers)

	if result.Error != nil {
		return foundUsers, result.Error
	}
	return foundUsers, nil
}

func (repo *UserRepository) Read(user *model.User) ([]model.User, error) {
	var foundUsers []model.User

	result := repo.DB.Where(user).Find(&foundUsers)

	if result.Error != nil {
		return foundUsers, result.Error
	}
	return foundUsers, nil
}

func (repo *UserRepository) Create(user *model.User) (*model.User, error) {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) SetAvatar(user *model.User) (*model.User, error) {
	var foundUser []model.User

	// check user
	resultUser := repo.DB.Limit(1).Where(&model.User{UserID: user.UserID}).Find(&foundUser)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}
	if len(foundUser) == 0 {
		return nil, errors.New("userId not found")
	}

	// set avatar
	foundUser[0].Avatar = user.Avatar

	// Update changes
	result := repo.DB.Save(&foundUser[0])
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser[0], nil
}
