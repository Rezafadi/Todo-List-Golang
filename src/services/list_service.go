package services

import (
	"TechnikalTes/helpers"
	"TechnikalTes/src/models"
	"TechnikalTes/src/repositories"
	"fmt"

	"gorm.io/gorm"
)

type listService struct {
	listRepository repositories.ListRepository
}

// Create implements ListService.
func (service *listService) Create(lists models.List) helpers.Response {
	var response helpers.Response

	if err := service.listRepository.Create(lists); err != nil {
			response.Status = 500
			response.Message = "Failed to create list." + err.Error()
	} else {
			response.Status = 200
			response.Message = "Success to create list."
	}

	return response
}

// Delete implements ListService.
func (service *listService) Delete(id int) helpers.Response {
	var response helpers.Response
	if err := service.listRepository.Delete(id); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete list : ", id)
	} else {
		response.Status = 200
		response.Message = fmt.Sprint("Success to delete list : ", id)
	}
	return response
}

// GetAll implements ListService.
func (service *listService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.listRepository.GetAll()
	if err != nil {
		response.Status = 500
		response.Message = "Failed to get lists." + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success to get lists."
		response.Data = data
	}
	return response
}

// GetById implements ListService.
func (service *listService) GetById(id int) helpers.Response {
	var response helpers.Response
	data, err := service.listRepository.GetById(id)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to get list : ", id , err.Error())
	} else {
		response.Status = 200
		response.Message = "Success to get list by Id."
		response.Data = data
	}
	return response
}

// Update implements ListService.
func (service *listService) Update(id int, lists models.List) helpers.Response {
	var response helpers.Response
	if err := service.listRepository.Update(id, lists); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to Update list : ", id)
	} else {
		response.Status = 200
		response.Message = "Success to Update list."
	}
	return response
}

type ListService interface {
	Create(lists models.List) helpers.Response
	Update(id int, lists models.List) helpers.Response
	Delete(id int) helpers.Response
	GetById(id int) helpers.Response
	GetAll() helpers.Response
}

func NewListService(db *gorm.DB) ListService {
	return &listService{
		listRepository: repositories.NewListRepository(db),
	}
}
