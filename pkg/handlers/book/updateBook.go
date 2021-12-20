package book

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func UpdateBook(id uint, mark string) {
	var book models.Book

	db := handlers.Database()
	db.First(&book, id)

	book.Mark = mark
	book.Number_Mark += 1

	db.Save(&book)
	db.Close()
}
