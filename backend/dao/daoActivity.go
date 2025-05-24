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

func DaoToDto(activity Activity) domain.Activity {
	return domain.Activity{
		Id:          activity.Id,
		Name:        activity.Name,
		Description: activity.Description,
		Capacity:    activity.Capacity,
		Category:    activity.Category,
		Profesor:    activity.Profesor,
		Day:         activity.Day,
		Hour:        activity.Hour,
	}
}
func GetActivities(DB *gorm.DB) ([]domain.Activity, error) {
	var activitiesDAO []Activity
	txn := DB.Find(&activitiesDAO)
	if txn.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", txn.Error)
	}
	var activities []domain.Activity
	for _, activity := range activitiesDAO {
		activities = append(activities, DaoToDto(activity))
	}
	return activities, nil
}
func GetActivityByID(DB *gorm.DB, id int) (domain.Activity, error) {
	var activityDAO Activity
	txn := DB.First(&activityDAO, "id = ?", id)
	if txn.Error != nil {
		return domain.Activity{}, fmt.Errorf("error getting activity: %w", txn.Error)
	}
	return DaoToDto(activityDAO), nil
}
func GetActivitiesByFilters(DB *gorm.DB, filters map[string]string) ([]domain.Activity, error) {
	var activitiesDAO []Activity
	query := DB
	// Palabra clave: busca en name, descriptionr
	if v := filters["keyword"]; v != "" {
		like := "%" + v + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", like, like)

	}
	if v := filters["category"]; v != "" {
		query = query.Where("category LIKE ?", "%"+v+"%")
	}
	if v := filters["hour"]; v != "" {
		query = query.Where("hour LIKE ?", "%"+v+"%")
	}
	txn := query.Find(&activitiesDAO)
	if txn.Error != nil {
		return nil, fmt.Errorf("error getting activities: %w", txn.Error)
	}
	var activities []domain.Activity
	for _, activity := range activitiesDAO {
		activities = append(activities, DaoToDto(activity))
	}
	if len(activitiesDAO) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return activities, nil
}
