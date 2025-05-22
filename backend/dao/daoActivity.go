package dao

import (
	"fmt"
	"gym-api/backend/domain"

	"gorm.io/gorm"
)

type Activity struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(350);not null"`
	Description string `gorm:"type:varchar(350);not null"`
	Capacity    int    `gorm:"not null"`
	Category    string `gorm:"type:varchar(350);not null"`
	Profesor    string `gorm:"type:varchar(350);not null"`
	Day         string `gorm:"type:varchar(350);not null"`
	Hour        string `gorm:"type:varchar(350);not null"`
}
type Activities []Activity

func GetActivities(DB *gorm.DB) ([]domain.Activity, error) {
	var activitiesDAO []Activity
	txn := DB.Find(&activitiesDAO)
	if txn.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", txn.Error)
	}
	var activities []domain.Activity
	for _, activity := range activitiesDAO {
		activities = append(activities, domain.Activity{
			Id:          activity.Id,
			Name:        activity.Name,
			Description: activity.Description,
			Capacity:    activity.Capacity,
			Category:    activity.Category,
			Profesor:    activity.Profesor,
			Day:         activity.Day,
			Hour:        activity.Hour,
		})
	}
	return activities, nil
}

func GetActivityByID(DB *gorm.DB, id int) (domain.Activity, error) {
	var activityDAO Activity
	txn := DB.First(&activityDAO, "id = ?", id)
	if txn.Error != nil {
		return domain.Activity{}, fmt.Errorf("error getting activity: %w", txn.Error)
	}
	return domain.Activity{
		Id:          activityDAO.Id,
		Name:        activityDAO.Name,
		Description: activityDAO.Description,
		Capacity:    activityDAO.Capacity,
		Category:    activityDAO.Category,
		Profesor:    activityDAO.Profesor,
		Day:         activityDAO.Day,
		Hour:        activityDAO.Hour,
	}, nil
}
