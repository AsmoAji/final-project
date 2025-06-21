package repository

import (
	"final-project/entity"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	FindAll() ([]entity.Activity, error)
	FindByID(id uint) (entity.Activity, error)
	Create(activity *entity.Activity) error
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) FindAll() ([]entity.Activity, error) {
	var activities []entity.Activity
	err := r.db.Find(&activities).Error
	return activities, err
}

func (r *activityRepository) FindByID(id uint) (entity.Activity, error) {
	var activity entity.Activity
	err := r.db.First(&activity, id).Error
	return activity, err
}

func (r *activityRepository) Create(activity *entity.Activity) error {
	return r.db.Create(&activity).Error
}
