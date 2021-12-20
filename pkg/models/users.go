package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	LastName  string `json:"name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rool      string `json:"rool"`
	Books     string `json:"books"`
}

type TransformedUser struct {
	Id        uint   `json:"id"`
	LastName  string `json:"name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rool      string `json:"rool"`
	Books     string `json:"books"`
}
