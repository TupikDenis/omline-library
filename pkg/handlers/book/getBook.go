package book

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func GetAllNameBooks(name string) []models.TransformedBook {
	var books []models.Book
	var _books []models.TransformedBook

	db := handlers.Database()
	db.Find(&books)
	db.Close()

	for _, item := range books {
		if item.Name == name {
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
	}

	return _books
}

func GetBookById(id uint) models.TransformedBook {
	var books []models.Book
	var _books []models.TransformedBook
	var book models.TransformedBook

	db := handlers.Database()
	db.Find(&books)
	db.Close()

	for _, item := range books {
		if item.ID == id {
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
	}

	if len(_books) == 1 {
		book = _books[0]
	}

	return book
}
