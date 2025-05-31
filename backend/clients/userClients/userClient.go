package clients

import (
	"errors"
	"fmt"
	"gym-api/backend/dao"

	"gorm.io/gorm"
)

// “Un repositorio no es la base de datos, es la puerta de entrada a ella desde el dominio.”
type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	GetUserByEmail(email string) (dao.User, error)
	GetUserByID(id int) (dao.User, error)
	InsertUser(name string,
		lastName string,
		email string,
		password string,
		birthDate string,
		sex string) (dao.User, error)
	EmailAlreadyExists(email string) error
}

func (mySQLDatasource UserRepository) GetUserByEmail(email string) (dao.User, error) {
	var user dao.User
	result := mySQLDatasource.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("%w", result.Error)
	}
	return user, nil
}
func (mySQLDatasource UserRepository) EmailAlreadyExists(email string) error {
	var user dao.User
	result := mySQLDatasource.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil
	}
	return errors.New("El email que ingresaste ya esta en uso")
}

func (mySQLDatasource UserRepository) GetUserByID(id int) (dao.User, error) {
	var user dao.User
	result := mySQLDatasource.DB.First(&user, id)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return user, nil
}
func (mySQLDatasource UserRepository) InsertUser(name string, lastName string, email string, password string, birthDate string, sex string) (dao.User, error) {
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
		return dao.User{}, fmt.Errorf("error inserting user: %w", txn.Error)
	}
	return user, nil
}
