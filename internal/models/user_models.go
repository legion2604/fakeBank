package models

import "time"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
}

type GetUser struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
}

type UserJson struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type LoginJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
