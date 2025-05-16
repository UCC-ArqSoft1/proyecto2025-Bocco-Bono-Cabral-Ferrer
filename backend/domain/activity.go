package domain

type Activity struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(350);not null" json:"name"`
	Description string `gorm:"type:varchar(350);not null" json:"description"`
	Capacity    int    `gorm:"not null" json:"capacity"`
	Category    string `gorm:"type:varchar(350);not null" json:"category"`
	Profesor    string `gorm:"type:varchar(350);not null" json:"profesor"`
	Day         string `gorm:"type:varchar(350);not null" json:"day"`
	Hour        string `gorm:"type:varchar(350);not null" json:"hour"`
}

type Activities []Activity
