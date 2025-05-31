package dao

import (
	"gym-api/backend/domain"
)

type Activity struct {
	Id          int                `gorm:"primaryKey"`
	Name        string             `gorm:"type:varchar(350);not null"`
	Description string             `gorm:"type:varchar(350);not null"`
	Capacity    int                `gorm:"not null"`
	Category    string             `gorm:"type:varchar(350);not null"`
	Profesor    string             `gorm:"type:varchar(350);not null"`
	Schedules   []ActivitySchedule `gorm:"foreignKey:ActivityId;"`
}

type ActivitySchedule struct {
	Id         int    `gorm:"primaryKey"`
	ActivityId int    `gorm:"not null"`
	Day        string `gorm:"type:varchar(10);not null;default:''"`
	StartTime  string `gorm:"type:varchar(5);not null;default:''"`
	EndTime    string `gorm:"type:varchar(5);not null;default:''"`
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
		Schedules:   ConvertActivitySchedules(activity.Schedules),
	}

}
func ConvertActivitySchedules(schedules []ActivitySchedule) []domain.ActivitySchedule {
	var domainSchedules []domain.ActivitySchedule
	for _, schedule := range schedules {
		domainSchedules = append(domainSchedules, domain.ActivitySchedule{
			Id:         schedule.Id,
			ActivityId: schedule.ActivityId,
			Day:        schedule.Day,
			StartTime:  schedule.StartTime,
			EndTime:    schedule.EndTime,
		})
	}
	return domainSchedules
}
