package domain

type User struct {
	Id         int      `gorm:"primaryKey" json:"id"`
	FirstName  string   `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName   string   `gorm:"type:varchar(100);not null" json:"last_name"`
	Email      string   `gorm:"type:varchar(250);not null;unique" json:"email"`
	Password   string   `gorm:"type:varchar(250);not null" json:"password"`
	Birth_date string   `gorm:"type:varchar(50);not null" json:"birth_date"`
	Sex        string   `gorm:"type:varchar(50);not null" json:"sex"`
	UserTypeId int      `json:"user_type_id"`
	UserType   UserType `gorm:"foreignKey:UserTypeId" json:"user_type"`
}

type Users []User
