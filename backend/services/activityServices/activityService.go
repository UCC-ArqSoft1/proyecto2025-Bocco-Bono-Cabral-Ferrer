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
	CreateActivity(name string, description string, capacity int, category string, profesor string, schedules []domain.ActivitySchedule) error
	DeleteActivity(id int) error
	UpdateActivity(id int, name string, description string, capacity int, category string, profesor string, schedules []domain.ActivitySchedule) error
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

func (a ActivityServiceImpl) CreateActivity(
	name string, description string, capacity int,
	category string, profesor string, schedules []domain.ActivitySchedule,
) error {
	validSchedules := []dao.ActivitySchedule{}
	for _, s := range schedules {
		if s.Day != "" && s.StartTime != "" && s.EndTime != "" {
			validSchedules = append(validSchedules, dao.ActivitySchedule{
				Day:       s.Day,
				StartTime: s.StartTime,
				EndTime:   s.EndTime,
			})
		}
	}
	return a.Repo.CreateActivity(name, description, capacity, category, profesor, validSchedules)
}
func (a ActivityServiceImpl) DeleteActivity(id int) error {
	daoActivity, err := a.Repo.GetActivityByID(id)
	if err != nil {
		return err
	}
	return a.Repo.DeleteActivity(daoActivity.Id)
}
func (a ActivityServiceImpl) UpdateActivity(
	id int, name string, description string, capacity int,
	category string, profesor string, schedules []domain.ActivitySchedule) error {
	_, err := a.Repo.GetActivityByID(id)
	if err != nil {
		return err
	}
	validSchedules := []dao.ActivitySchedule{}
	for _, s := range schedules {
		if s.Day != "" && s.StartTime != "" && s.EndTime != "" {
			validSchedules = append(validSchedules, dao.ActivitySchedule{
				Day:       s.Day,
				StartTime: s.StartTime,
				EndTime:   s.EndTime,
			})
		}
	}
	return a.Repo.UpdateActivity(id, name, description, capacity, category, profesor, validSchedules)
}
