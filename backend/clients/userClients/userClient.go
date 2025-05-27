package clients

import (
	"fmt"
	"gym-api/backend/dao"

	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	DB *gorm.DB
}

type UserClient interface {
	GetUsers() ([]dao.User, error)
	GetUserByID(id int) (dao.User, error)
	GetUsersByFilters(keyword string) ([]dao.User, error)
}

func (mySQLDatasource MySQLUserRepository) GetUserByEmail(email string) (dao.User, error) {
	var user dao.User
	result := mySQLDatasource.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return user, nil
}

func (mySQLDatasource MySQLUserRepository) InsertUser(name string, lastName string, email string, password string, birthDate string, sex string) (int, error) {
	user := dao.User{
		FirstName:    name,
		LastName:     lastName,
		Email:        email,
		PasswordHash: password,
		Birth_date:   birthDate,
		Sex:          sex,
		UserTypeId:   2,
	}
	txn := mySQLDatasource.DB.Create(&user)
	if txn.Error != nil {
		return 0, fmt.Errorf("error inserting user: %w", txn.Error)
	}
	return user.Id, nil
}
