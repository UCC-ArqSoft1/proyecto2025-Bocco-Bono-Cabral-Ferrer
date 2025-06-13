package dao

type UserType struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}
type UserTypes []UserType
