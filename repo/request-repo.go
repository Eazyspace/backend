package repo

import (
	"errors"

	"github.com/Eazyspace/model"
	"gorm.io/gorm"
)

type IRequestRepository interface {
	Create(lastname, firstname string) (*model.Request, error)
	ReadByID(id int) (*model.Request, error)
	ReadAll() (*model.Request, error)
	Update(id int, lastname, firstname string) (*model.Request, error)
	Delete(id int) error
}

type RequestRepository struct {
	DB *gorm.DB
}

func NewRequestRepo(db *gorm.DB) *RequestRepository {
	return &RequestRepository{
		DB: db,
	}
}

func (repo *RequestRepository) Read(Request *model.Request) ([]model.Request, error) {
	var foundRequest []model.Request

	// select * from Request where Request.Request_code = Request.Request_code
	result := repo.DB.Where(Request).Find(&foundRequest)

	if result.Error != nil {
		return foundRequest, result.Error
	}
	return foundRequest, nil
}

func (repo *RequestRepository) Create(request *model.Request) (*model.Request, error) {
	var foundRoom []model.Room

	resultRoom := repo.DB.Limit(1).Where(&model.Room{RoomID: request.RoomID}).Find(&foundRoom)

	if resultRoom.Error != nil {
		return nil, resultRoom.Error
	}
	if len(foundRoom) == 0 {
		return nil, errors.New("roomId not found")
	}

	/*result := repo.DB.Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}*/
	return request, nil
}
