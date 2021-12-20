package user

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func DeleteUser(user *models.User, id uint) {
	db := handlers.Database()
	db.First(&user, id)

	if id <= 0 {
		return
	}

	db.Delete(&user)
}
