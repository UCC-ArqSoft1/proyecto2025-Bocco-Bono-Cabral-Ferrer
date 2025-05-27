package services

import (
	clients "gym-api/backend/clients/activityclient"
	"gym-api/backend/dao"
	"gym-api/backend/domain"
)

type ActivityServiceImpl struct {
	Repo clients.ActivityRepositoryInterface
}
type ActivityServiceInterface interface {
	GetActivities() ([]domain.Activity, error)
	GetactivityByID(id int) (domain.Activity, error)
	GetActivitiesByFilters(keyword string) ([]domain.Activity, error)
}

func (a ActivityServiceImpl) GetActivities() ([]domain.Activity, error) {
	daoActivities, err := a.Repo.GetActivities()
	if err != nil {
		return nil, err
	}
	var dtoActivities []domain.Activity
	for _, activity := range daoActivities {
		dtoActivities = append(dtoActivities, dao.DaoToDto(activity))
	}
	return dtoActivities, nil
}
func (a ActivityServiceImpl) GetactivityByID(id int) (domain.Activity, error) {
	daoActivity, err := a.Repo.GetActivityByID(id)
	if err != nil {
		return domain.Activity{}, err
	}
	var dtoActivity domain.Activity
	dtoActivity = dao.DaoToDto(daoActivity)
	return dtoActivity, nil
}
func (a ActivityServiceImpl) GetActivitiesByFilters(keyword string) ([]domain.Activity, error) {
	daoActivities, err := a.Repo.GetActivitiesByFilters(keyword)
	if err != nil {
		return nil, err
	}
	var dtoActivities []domain.Activity
	for _, activity := range daoActivities {
		dtoActivities = append(dtoActivities, dao.DaoToDto(activity))
	}
	return dtoActivities, nil
}
