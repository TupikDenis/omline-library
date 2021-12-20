package book

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func DeleteBook(book *models.Book, id uint) {
	db := handlers.Database()
	db.First(&book, id)

	if id <= 0 {
		return
	}

	db.Delete(&book)
}
