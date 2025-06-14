package repository

import (
	"final-project/entity"

	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() ([]entity.Item, error)
	FindByID(id uint) (entity.Item, error)
	Create(item entity.Item) error
	Update(item entity.Item) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) FindAll() ([]entity.Item, error) {
	var items []entity.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id uint) (entity.Item, error) {
	var item entity.Item
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *itemRepository) Create(item entity.Item) error {
	return r.db.Create(&item).Error
}

func (r *itemRepository) Update(item entity.Item) error {
	return r.db.Save(&item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Item{}, id).Error
}
