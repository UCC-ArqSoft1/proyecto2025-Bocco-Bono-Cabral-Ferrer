package dao

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
