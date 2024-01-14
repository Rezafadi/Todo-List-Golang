package services

import (
	"TechnikalTes/helpers"
	"TechnikalTes/src/models"
	"TechnikalTes/src/repositories"
	"fmt"

	"gorm.io/gorm"
)

type subListService struct {
	subListRepository repositories.SubListRepository
}

// Create implements SubListService.
func (service *subListService) Create(subLists models.SubList) helpers.Response {
	var response helpers.Response
	if err := service.subListRepository.Create(subLists); err != nil {
		response.Status = 500
		response.Message = "Failed to Create Sub List" + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success to Create Sub List"
	}
	return response
}

// Delete implements SubListService.
func (service *subListService) Delete(id int) helpers.Response {
	var response helpers.Response
	if err := service.subListRepository.Delete(id); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete Sub List : ", id)
	} else {
		response.Status = 200
		response.Message = fmt.Sprint("Success to delete Sub List : ", id)
	}
	return response
}

// GetAll implements SubListService.
func (service *subListService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.subListRepository.GetAll()
	if err != nil {
		response.Status = 500
		response.Message = "Failed to get Sub Lists" + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success to get Sub Lists"
		response.Data = data
	}
	return response
}

// GetById implements SubListService.
func (service *subListService) GetById(id int) helpers.Response {
	var response helpers.Response
	data, err := service.subListRepository.GetById(id)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to get Sub List : ", id, err.Error())
	} else {
		response.Status = 200
		response.Message = "Success to get Sub List by Id"
		response.Data = data
	}
	return response
}

// Update implements SubListService.
func (service *subListService) Update(id int, subLists models.SubList) helpers.Response {
	var response helpers.Response
	if err := service.subListRepository.Update(id, subLists); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to Update Sub List : ", id, err.Error())
	} else {
		response.Status = 200
		response.Message = "Success to Update Sub List"
	}
	return response
}

type SubListService interface {
	Create(subLists models.SubList) helpers.Response
	Update(id int, subLists models.SubList) helpers.Response
	Delete(id int) helpers.Response
	GetById(id int) helpers.Response
	GetAll() helpers.Response
}

func NewSubListService(db *gorm.DB) SubListService {
	return &subListService{
		subListRepository: repositories.NewSubListRepository(db),
	}
}
