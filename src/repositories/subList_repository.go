package repositories

import (
	"TechnikalTes/src/models"

	"gorm.io/gorm"
)

type dbSubList struct {
	Conn *gorm.DB
}

// Create implements SubListRepository.
func (db *dbSubList) Create(subLists models.SubList) error {
	return db.Conn.Create(&subLists).Error
}

// Delete implements SubListRepository.
func (db *dbSubList) Delete(id int) error {
	return db.Conn.Delete(&models.SubList{}, id).Error
}

// GetAll implements SubListRepository.
func (db *dbSubList) GetAll() ([]models.SubList, error) {
	var data []models.SubList
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements SubListRepository.
func (db *dbSubList) GetById(id int) (models.SubList, error) {
	var data models.SubList
	result := db.Conn.Where("id = ?", id).First(&data)
	return data, result.Error
}

// Update implements SubListRepository.
func (db *dbSubList) Update(id int, subLists models.SubList) error {
	return db.Conn.Table("sub_lists").Where("id = ?", id).Omit("id").Updates(subLists).Error
}


type SubListRepository interface {
	Create(subLists models.SubList) error
	Update(id int, subLists models.SubList) error
	Delete(id int) error
	GetById(id int) (models.SubList, error)
	GetAll() ([]models.SubList, error)
}

func NewSubListRepository(Conn *gorm.DB) SubListRepository {
	return &dbSubList{Conn: Conn}
}
