package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id           int    `gorm:"primaryKey"`
	FirstName    string `gorm:"type:varchar(100);not null"`
	LastName     string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(250);not null;unique"`
	PasswordHash string `gorm:"type:varchar(250);not null"`
	Birth_date   string `gorm:"type:varchar(50);not null"`
	Sex          string `gorm:"type:varchar(50);not null"`
	UserTypeId   int
	UserType     UserType `gorm:"foreignKey:UserTypeId"`
}
type Users []User

/*
	func GetUserByEmail(DB *gorm.DB, email string) (User, error) {
		var userDAO User
		txn := DB.First(&userDAO, "email = ?", email)
		if txn.Error != nil {
			return User{}, fmt.Errorf("error getting user: %w", txn.Error)
		}
		return userDAO, nil
	}
*/
func InsertUser(DB *gorm.DB, name string, lastName string, email string, password string, birthDate string, sex string) (int, error) {
	userDAO := User{
		FirstName:    name,
		LastName:     lastName,
		Email:        email,
		PasswordHash: password,
		Birth_date:   birthDate,
		Sex:          sex,
		UserTypeId:   1,
	}
	txn := DB.Create(&userDAO)
	if txn.Error != nil {
		return 0, fmt.Errorf("error inserting user: %w", txn.Error)
	}
	return userDAO.Id, nil
}
