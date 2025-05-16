package domain

type UserType struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}

type UserTypes []UserType
