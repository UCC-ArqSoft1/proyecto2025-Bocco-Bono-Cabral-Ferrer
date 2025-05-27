package services

import (
	clients "gym-api/backend/clients/activityclient"
	"gym-api/backend/dao"
	"gym-api/backend/domain"
)

type ActivityServices struct {
	ActivityClients clients.MySQLActivityRepository
}
type ActivityService interface {
	GetActivities() ([]domain.Activity, error)
	GetactivityByID(id int) (domain.Activity, error)
	GetActivitiesByFilters(keyword string) ([]domain.Activity, error)
}

func (a ActivityServices) GetActivities() ([]domain.Activity, error) {
	daoActivities, err := a.ActivityClients.GetActivities()
	if err != nil {
		return nil, err
	}
	var dtoActivities []domain.Activity
	for _, activity := range daoActivities {
		dtoActivities = append(dtoActivities, dao.DaoToDto(activity))
	}
	return dtoActivities, nil
}
func (a ActivityServices) GetactivityByID(id int) (domain.Activity, error) {
	daoActivity, err := a.ActivityClients.GetActivityByID(id)
	if err != nil {
		return domain.Activity{}, err
	}
	var dtoActivity domain.Activity
	dtoActivity = dao.DaoToDto(daoActivity)
	return dtoActivity, nil
}
func (a ActivityServices) GetActivitiesByFilters(keyword string) ([]domain.Activity, error) {
	daoActivities, err := a.ActivityClients.GetActivitiesByFilters(keyword)
	if err != nil {
		return nil, err
	}
	var dtoActivities []domain.Activity
	for _, activity := range daoActivities {
		dtoActivities = append(dtoActivities, dao.DaoToDto(activity))
	}
	return dtoActivities, nil
}
