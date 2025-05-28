package clients

import (
	"errors"
	"gym-api/backend/dao"

	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	DB *gorm.DB
}

type EnrollmentRepositoryInterface interface {
	IsEnrolled(userId, activityId int) (bool, error)
	CountEnrollments(activityId int) (int, error)
	CreateEnrollment(userId, activityId int) error
}

func (er EnrollmentRepository) IsEnrolled(userId, activityId int) (bool, error) {
	var enrollments dao.Enrollment
	result := er.DB.First(&enrollments, userId, activityId)
	if result.Error != nil {
		return true, result.Error
	}
	if enrollments.Id != 0 {
		return true, nil
	}
	return false, nil
}

func (er EnrollmentRepository) CountEnrollments(activityId int) (int, error) {
	var counter int
	result := er.DB.Raw("Select COUNT(enrollments.id) FROM enrollments INNER JOIN activities ON enrollments.activity_id = activities.id").
		Scan(&counter)
	if result.Error != nil {
		return 0, result.Error
	}
	return counter, nil
}

func (er EnrollmentRepository) CreateEnrollment(userId, activityId int) error {
	isenrolled, err := er.IsEnrolled(userId, activityId)
	if err != nil {
		return err
	}
	if isenrolled != false {
		return errors.New("Ya estas inscripto en esta actividad")
	}
}
