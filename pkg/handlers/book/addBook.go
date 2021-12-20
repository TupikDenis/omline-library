package book

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"

	"github.com/jinzhu/gorm"
)

func AddBook(name string, author string, fileName string) {
	book := models.Book{
		Model:       gorm.Model{},
		Name:        name,
		Author:      author,
		FileName:    fileName,
		Mark:        "0.0",
		Number_Mark: 0,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.Book{})
	db.Create(&book)
	db.Close()
}
