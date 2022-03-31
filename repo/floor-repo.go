package repo

import (
	"github.com/Eazyspace/model"
	"gorm.io/gorm"
)

type IFloorRepository interface {
	Create(lastname, firstname string) (*model.Floor, error)
	ReadByID(id int) (*model.Floor, error)
	ReadAll() (*model.Floor, error)
	Update(id int, lastname, firstname string) (*model.Floor, error)
	Delete(id int) error
}

type FloorRepository struct {
	DB *gorm.DB
}

func NewFloorRepo(db *gorm.DB) *FloorRepository {
	return &FloorRepository{
		DB: db,
	}
}

func (repo *FloorRepository) ReadAll() ([]model.Floor, error) {
	var foundFloors []model.Floor

	result := repo.DB.Find(&foundFloors)

	if result.Error != nil {
		return foundFloors, result.Error
	}
	return foundFloors, nil
}

func (repo *FloorRepository) Read(floor *model.Floor) ([]model.Floor, error) {
	var foundFloors []model.Floor

	result := repo.DB.Where(floor).Find(&foundFloors)

	if result.Error != nil {
		return foundFloors, result.Error
	}
	return foundFloors, nil
}

func (repo *FloorRepository) Create(floor *model.Floor) (*model.Floor, error) {
	result := repo.DB.Create(&floor)
	if result.Error != nil {
		return nil, result.Error
	}
	return floor, nil
}
