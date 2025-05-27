package clients

import (
	"gym-api/backend/dao"

	"gorm.io/gorm"
)

type MySQLActivityRepository struct {
	DB *gorm.DB
}

type ActivityClients interface {
	GetActivities() ([]dao.Activity, error)
	GetActivityByID(id int) (dao.Activity, error)
	GetActivitiesByFilters(keyword string) ([]dao.Activity, error)
}

// service
func (mySQLDatasource MySQLActivityRepository) GetActivities() ([]dao.Activity, error) {
	var activities []dao.Activity
	result := mySQLDatasource.DB.Preload("Schedules").Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (mySQLDatasource MySQLActivityRepository) GetActivityByID(id int) (dao.Activity, error) {
	var activity dao.Activity

	result := mySQLDatasource.DB.First(&activity, id)
	if result.Error != nil {
		return dao.Activity{}, result.Error
	}
	return activity, nil
}

func (mySQLDatasource MySQLActivityRepository) GetActivitiesByFilters(keyword string) ([]dao.Activity, error) {
	var activities []dao.Activity
	Keyword := "%" + keyword + "%"
	result := mySQLDatasource.DB.
		Joins("JOIN activity_schedules ON activity_schedules.activity_id = activities.id").
		Where(`
		activities.name LIKE ? OR 
		activities.description LIKE ? OR 
		activities.category LIKE ? OR 
		activity_schedules.day LIKE ? OR 
		activity_schedules.start_time LIKE ?
	`, Keyword, Keyword, Keyword, Keyword, Keyword).
		Preload("Schedules").
		Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}
