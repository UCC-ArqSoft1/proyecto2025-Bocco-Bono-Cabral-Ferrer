package domain

import "time"

type Enrollment struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	UserId         int       `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId" json:"user"`
	ActivityId     int       `json:"activity_id"`
	Activity       Activity  `gorm:"foreignKey:ActivityId" json:"activity"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}

type Enrollments []Enrollment
