package user

import (
	"log"
	"myLibrary/pkg/handlers"
	"myLibrary/pkg/models"
)

func GetUserById(id uint) models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser
	var user models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	db.Close()

	for _, item := range users {
		if item.ID == id {
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
	}

	if len(_users) == 1 {
		user = _users[0]
	}

	return user
}

func Login(user_password string, user_email string) models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser
	var user models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	db.Close()

	for _, item := range users {
		if item.Email == user_email && item.Password == user_password {
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
	}

	if len(_users) == 1 {
		user = _users[0]
	}

	log.Println(user)

	return user
}
