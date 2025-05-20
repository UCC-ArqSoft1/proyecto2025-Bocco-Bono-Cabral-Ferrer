package domain

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Birth_date string `json:"birth_date"`
	Sex        string `json:"sex"`
}

type Users []User
