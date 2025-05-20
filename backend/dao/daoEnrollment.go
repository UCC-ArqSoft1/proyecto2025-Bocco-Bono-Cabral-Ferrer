package dao

import "time"

type Enrollment struct {
	Id             int `gorm:"primaryKey"`
	UserId         int
	User           User `gorm:"foreignKey:UserId"`
	ActivityId     int
	Activity       Activity `gorm:"foreignKey:ActivityId"`
	EnrollmentDate time.Time
}

type Enrollments []Enrollment
