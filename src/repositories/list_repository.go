package repositories

import (
	"TechnikalTes/src/models"

	"gorm.io/gorm"
)

type dbList struct {
	Conn *gorm.DB
}

// Create implements ListRepository.
func (db *dbList) Create(lists models.List) error {
	return db.Conn.Create(&lists).Error
}

// Delete implements ListRepository.
func (db *dbList) Delete(id int) error {
	return db.Conn.Delete(&models.List{}, id).Error
}

// GetAll implements ListRepository.
func (db *dbList) GetAll() ([]models.List, error) {
	var lists []models.List
    if err := db.Conn.Preload("SubLists").Find(&lists).Error; err != nil {
        return nil, err
    }
    return lists, nil
}

// GetById implements ListRepository.
func (db *dbList) GetById(id int) (models.List, error) {
	var data models.List

	result := db.Conn.Preload("SubLists").Where("id = ?", id).First(&data)
	return data, result.Error
}

// Update implements ListRepository.
func (db *dbList) Update(id int, lists models.List) error {
	return db.Conn.Where("id = ?", id).Updates(lists).Error
}

type ListRepository interface {
	Create(lists models.List) error
	Update(id int, lists models.List) error
	Delete(id int) error
	GetById(id int) (models.List, error)
	GetAll() ([]models.List, error)
}

func NewListRepository(Conn *gorm.DB) ListRepository {
	return &dbList{Conn: Conn}
}
