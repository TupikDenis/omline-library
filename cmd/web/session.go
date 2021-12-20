package main

import (
	"myLibrary/pkg/handlers/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func userSession(c *gin.Context) (uint, string) {
	session := sessions.Default(c)
	var id uint
	var fullName string

	i := session.Get("id")

	if i == nil {
		id = 0
	} else {
		id = session.Get("id").(uint)
	}

	session.Set("id", id)

	userInfo := user.GetUserById(id)
	fullName = userInfo.LastName + " " + userInfo.FirstName

	session.Set("full_name", fullName)
	session.Set("email", userInfo.Email)
	session.Set("rool", userInfo.Rool)
	session.Save()

	return id, fullName
}
