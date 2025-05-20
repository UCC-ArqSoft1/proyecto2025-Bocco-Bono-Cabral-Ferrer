package dao

type Activity struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(350);not null"`
	Description string `gorm:"type:varchar(350);not null"`
	Capacity    int    `gorm:"not null" json:"capacity"`
	Category    string `gorm:"type:varchar(350);not null"`
	Profesor    string `gorm:"type:varchar(350);not null"`
	Day         string `gorm:"type:varchar(350);not null"`
	Hour        string `gorm:"type:varchar(350);not null"`
}
type Activities []Activity
