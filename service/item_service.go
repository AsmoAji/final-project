package service

import (
	"final-project/entity"
	"final-project/repository"
)

type ItemService interface {
	GetAllItems() ([]entity.Item, error)
	CreateItem(item *entity.Item) error
	UpdateItem(item *entity.Item) error
	DeleteItem(id uint) error
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) ItemService {
	return &itemService{r}
}

func (s *itemService) GetAllItems() ([]entity.Item, error) {
	return s.repo.FindAll()
}

func (s *itemService) CreateItem(item *entity.Item) error {
	return s.repo.Create(item)
}

func (s *itemService) UpdateItem(item *entity.Item) error {
	return s.repo.Update(item)
}

func (s *itemService) DeleteItem(id uint) error {
	return s.repo.Delete(id)
}
