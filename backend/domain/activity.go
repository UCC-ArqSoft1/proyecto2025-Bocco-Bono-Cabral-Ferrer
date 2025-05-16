package domain

type Activity struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(350);not null"`
	Description string `gorm:"type:varchar(350);not null"`
	Capacity    int    `gorm:"not null"`
	category    string `gorm:"type:varchar(350);not null"`
	profesor    string `gorm:"type:varchar(350);not null"`
	day         string `gorm:"type:varchar(350);not null"`
	hour        string `gorm:"type:varchar(350);not null"`
}
