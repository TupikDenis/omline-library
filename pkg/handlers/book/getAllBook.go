package book

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func GetAllBooks() []models.TransformedBook {
	var books []models.Book
	var _books []models.TransformedBook

	db := handlers.Database()
	db.Find(&books)
	for _, item := range books {
		_books = append(
			_books, models.TransformedBook{
				Id:          item.ID,
				Name:        item.Name,
				Author:      item.Author,
				FileName:    item.FileName,
				Mark:        item.Mark,
				Number_Mark: item.Number_Mark,
			})
	}

	return _books
}
