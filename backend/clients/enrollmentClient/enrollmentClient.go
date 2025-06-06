package clients

import (
	"errors"
	"gym-api/backend/dao"
	"time"

	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	DB *gorm.DB
}

type EnrollmentRepositoryInterface interface {
	IsEnrolled(userId, activityId int) (bool, error)
	CountEnrollmentsAndCapacity(activityId int) (int, int, error)
	CreateEnrollment(userId int, activityId int, date time.Time) error
	GetUserEnrollments(userId int) ([]dao.Enrollment, error)
}

func (er EnrollmentRepository) IsEnrolled(userId, activityId int) (bool, error) {
	var enrollment dao.Enrollment

	result := er.DB.
		Where("user_id = ? AND activity_id = ?", userId, activityId).
		First(&enrollment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (er EnrollmentRepository) CountEnrollmentsAndCapacity(activityId int) (int, int, error) {
	var counter, capacity int
	result := er.DB.Raw("SELECT COUNT(*) FROM enrollments WHERE activity_id = ?", activityId).
		Scan(&counter)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	result = er.DB.Raw("SELECT capacity FROM activities WHERE id = ?", activityId).
		Scan(&capacity)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return counter, capacity, nil
}

func (er EnrollmentRepository) CreateEnrollment(userId int, activityId int, date time.Time) error {
	enrollment := dao.Enrollment{
		UserId:         userId,
		ActivityId:     activityId,
		EnrollmentDate: date,
	}
	err := er.DB.Create(&enrollment)
	if err != nil {
		return err.Error
	}
	return nil
}

func (er EnrollmentRepository) GetUserEnrollments(userId int) ([]dao.Enrollment, error) {
	var enrollments []dao.Enrollment
	result := er.DB.
		Preload("Activity.Schedules").
		Where("user_id = ?", userId).
		Find(&enrollments)

	if result.Error != nil {
		return nil, result.Error
	}
	return enrollments, nil
}
