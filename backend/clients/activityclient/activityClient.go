package clients

import (
	"gym-api/backend/domain"

	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

type ActivityClients interface {
	GetActivities() ([]domain.Activity, error)
	GetActivityByID(id int) (domain.Activity, error)
	GetActivitiesByFilters(keyword string) ([]domain.Activity, error)
}

// service
func (mySQLDatasource MySQL) GetActivities() ([]domain.Activity, error) {
	var activities []domain.Activity
	result := mySQLDatasource.DB.Preload("activity_schedule").
		Joins("JOIN activity_schedule ON activity_schedule.activity_id = activities.id").
		Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}
func (mySQLDatasource MySQL) GetActivityByID(id int) (domain.Activity, error) {
	var activity domain.Activity

	result := mySQLDatasource.DB.First(&activity, id)
	if result.Error != nil {
		return domain.Activity{}, result.Error
	}
	return activity, nil
}

func (mySQLDatasource MySQL) GetActivitiesByFilters(keyword string) ([]domain.Activity, error) {
	var activities []domain.Activity
	Keyword := "%" + keyword + "%"
	result := mySQLDatasource.DB.Preload("activity_schedule").
		Joins("JOIN activity_schedule ON activity_schedule.activity_id = activities.id").
		Where("activities.name LIKE ? OR activities.description LIKE ? OR activities.category LIKE ? OR activity_schedule.day LIKE ? OR activity_schedule.start_time LIKE ?", Keyword, Keyword, Keyword, Keyword, Keyword).
		Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}
