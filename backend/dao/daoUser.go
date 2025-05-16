package dao

type User struct {
	Id         int    `json:"id"`
	First_name string `gorm:"type:varchar(100);not null" json:"first_name"`
	Last_name  string `gorm:"type:varchar(100);not null" json:"last_name"`
	Email      string `gorm:"type:varchar(250);not null;unique" json:"email"`
	Password   string `gorm:"type:varchar(250);not null" json:"password"`
	Birth_date string `gorm:"type:varchar(50);not null" json:"birth_date"`
	Sex        string `gorm:"type:varchar(50);not null" json:"sex"`
}
