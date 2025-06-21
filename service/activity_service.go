package service

import (
	"final-project/entity"
	"final-project/repository"
)

type ActivityService interface {
	GetAllActivities() ([]entity.Activity, error)
	GetActivityByID(id uint) (entity.Activity, error)
	CreateActivity(activity *entity.Activity) error
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(r repository.ActivityRepository) ActivityService {
	return &activityService{r}
}

func (s *activityService) GetAllActivities() ([]entity.Activity, error) {
	return s.repo.FindAll()
}

func (s *activityService) GetActivityByID(id uint) (entity.Activity, error) {
	return s.repo.FindByID(id)
}

func (s *activityService) CreateActivity(activity *entity.Activity) error {
	return s.repo.Create(activity)
}
