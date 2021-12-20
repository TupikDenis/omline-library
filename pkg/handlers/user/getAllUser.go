package user

import (
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func GetAllUsers() []models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	for _, item := range users {
		_users = append(
			_users, models.TransformedUser{
				Id:        item.ID,
				LastName:  item.LastName,
				FirstName: item.FirstName,
				Email:     item.Email,
				Password:  item.Password,
				Rool:      item.Rool,
				Books:     item.Books,
			})
	}

	return _users
}
