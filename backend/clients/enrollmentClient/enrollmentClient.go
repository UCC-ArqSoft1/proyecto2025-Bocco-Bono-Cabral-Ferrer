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

// Struct temporal con fecha como []uint8

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
	type EnrollmentSQL struct {
		UserId         int
		ActivityId     int
		EnrollmentDate []uint8
	}
	var rawEnrollments []EnrollmentSQL
	result := er.DB.
		Table("enrollments").
		Select("user_id, activity_id, enrollment_date").
		Where("user_id = ?", userId).
		Find(&rawEnrollments)

	if result.Error != nil {
		return nil, result.Error
	}

	var enrollments []dao.Enrollment

	for _, raw := range rawEnrollments {
		parsedDate, err := time.Parse("2006-01-02 15:04:05", string(raw.EnrollmentDate))
		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, dao.Enrollment{
			UserId:         raw.UserId,
			ActivityId:     raw.ActivityId,
			EnrollmentDate: parsedDate,
		})
	}

	return enrollments, nil
}
