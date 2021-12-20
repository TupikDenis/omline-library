package main

import (
	"log"
	"myLibrary/pkg/handlers/user"
	"myLibrary/pkg/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func addUserProcess(c *gin.Context) {
	lastName := c.PostForm("lastname")
	firstName := c.PostForm("firstname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user.AddUser(lastName, firstName, email, password, "user", "")
	c.Redirect(http.StatusFound, "/login")
}

func getAllUsersProcess(c *gin.Context) {

}

func loginuser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := user.Login(password, email)
	if user.Id < 0 {
		log.Println("Нет такого пользователя")
		return
	}

	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("full_name", user.FirstName+" "+user.LastName)
	session.Set("rool", user.Rool)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func updateUserPasswordProcess(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	password := c.PostForm("password")
	newPassword := c.PostForm("new_password")

	user.UpdateUserPassword(idu, password, newPassword)
}

func updateUserNameProcess(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	lastName := c.PostForm("lastname")
	firstName := c.PostForm("firstname")

	user.UpdateUserName(idu, lastName, firstName)
}

func deleteUserProcess(c *gin.Context) {
	var _user models.User
	id := c.Param("id")
	idu := convertUInt(id)

	user.DeleteUser(&_user, idu)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func profileExit(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
