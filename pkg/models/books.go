package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	FileName    string `json:"file_name"`
	Mark        string `json:"mark"`
	Number_Mark int    `json:"number_mark"`
}

type TransformedBook struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	FileName    string `json:"file_name"`
	Mark        string `json:"mark"`
	Number_Mark int    `json:"number_mark"`
}
