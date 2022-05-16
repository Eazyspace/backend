package repo

import (
	"errors"
	"time"

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

func (repo *RequestRepository) ReadWithFloorId(floorId int, Request *model.Request) ([]model.Request, error) {
	var foundRequest []model.Request

	// select * from Request where Request.Request_code = Request.Request_code
	result := repo.DB.Table("requests").Joins("JOIN rooms on rooms.room_id = requests.room_id").Where("rooms.floor_id = ?", floorId).Where(Request).Find(&foundRequest)

	if result.Error != nil {
		return foundRequest, result.Error
	}
	return foundRequest, nil
}

func (repo *RequestRepository) Create(request *model.Request) (*model.Request, error) {
	// check userId
	var foundUser []model.User
	resultUser := repo.DB.Limit(1).Where(&model.User{UserID: request.UserID}).Find(&foundUser)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}
	if len(foundUser) == 0 {
		return nil, errors.New("userId not found")
	}

	// check roomId
	var foundRoom []model.Room
	resultRoom := repo.DB.Limit(1).Where(&model.Room{RoomID: request.RoomID}).Find(&foundRoom)
	if resultRoom.Error != nil {
		return nil, resultRoom.Error
	}
	if len(foundRoom) == 0 {
		return nil, errors.New("roomId not found")
	}
	if foundRoom[0].Status != 1 {
		return nil, errors.New("current room is not available")
	}
	if foundRoom[0].MaxCapacity < request.NumberOfPeople {
		return nil, errors.New("room capacity is not enough")
	}

	// check request time
	if request.StartTime.Before(time.Now()) || request.StartTime.After(request.EndTime) || request.StartTime.Equal(request.EndTime) {
		return nil, errors.New("invalid startime and end time")
	}
	var foundRequest []model.Request
	resultRequest := repo.DB.Limit(1).Where("room_id = ? AND status = ? AND NOT (? < start_time OR end_time < ?)", request.RoomID, 2, request.EndTime, request.StartTime).Find(&foundRequest)
	if resultRequest.Error != nil {
		return nil, resultRequest.Error
	}
	if len(foundRequest) > 0 {
		return nil, errors.New("current room has been booked at that time")
	}

	// set request pending state
	request.Status = 1

	result := repo.DB.Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}
	return request, nil
}

func (repo *RequestRepository) UpdateStatus(Request *model.Request) (*model.Request, error) {
	var request *model.Request
	findRequest := repo.DB.Where("request_id = ?", Request.RequestID).First(&request)
	if findRequest.Error != nil {
		return nil, findRequest.Error
	}
	if request.Status != 1 {
		return nil, errors.New("status must be pending to be updated")
	}
	result := repo.DB.Model(&request).Updates(
		&model.Request{Status: Request.Status, ResponseNote: Request.ResponseNote})
	if result.Error != nil {
		return nil, result.Error
	}
	return Request, nil
}

func (repo *RequestRepository) CheckIn(Request *model.Request) (bool, error) {
	var result = repo.DB.Where(Request).Updates(&model.Request{Status: 4})
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
