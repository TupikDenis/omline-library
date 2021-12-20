package user

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func AddUser(lastName string, firstName string, email string, password string, rool string, books string) {
	user := models.User{
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  password,
		Rool:      rool,
		Books:     books,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.User{})
	db.Create(&user)
	db.Close()
}
