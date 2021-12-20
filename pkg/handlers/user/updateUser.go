package user

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func UpdateUserName(id uint, last string, first string) {
	var user models.User

	db := handlers.Database()
	db.First(&user, id)

	user.FirstName = first
	user.LastName = last

	db.Save(&user)
	db.Close()
}

func UpdateUserPassword(id uint, password string, newPassword string) {
	var user models.User

	db := handlers.Database()
	db.First(&user, id)

	if user.Password != password {
		return
	}

	user.Password = newPassword

	db.Save(&user)
	db.Close()
}
