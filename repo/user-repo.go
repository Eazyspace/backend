package repo

import (
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
