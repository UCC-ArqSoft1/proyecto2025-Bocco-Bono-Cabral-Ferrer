package services

import (
	clients "gym-api/backend/clients/activityclient"
	"gym-api/backend/domain"
)

type ActivityServices struct {
	ActivityClients clients.MySQL
}
type ActivityService interface {
	GetActivities() ([]domain.Activity, error)
	GetactivityByID(id int) (domain.Activity, error)
	GetActivitiesByFilters(keyword string) ([]domain.Activity, error)
}

func (a ActivityServices) GetActivities() ([]domain.Activity, error) {
	return a.ActivityClients.GetActivities()
}
func (a ActivityServices) GetactivityByID(id int) (domain.Activity, error) {
	return a.ActivityClients.GetActivityByID(id)
}
func (a ActivityServices) GetActivitiesByFilters(keyword string) ([]domain.Activity, error) {
	return a.ActivityClients.GetActivitiesByFilters(keyword)
}
