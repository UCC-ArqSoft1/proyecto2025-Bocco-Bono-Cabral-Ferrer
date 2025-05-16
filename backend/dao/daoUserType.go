package dao

type UserType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserTypes []UserType
