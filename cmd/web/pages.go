package main

import (
	"myLibrary/pkg/handlers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	id, fullName := userSession(c)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"name":      "Главная страница",
		"id":        id,
		"full_name": fullName,
	})
}

func login(c *gin.Context) {
	id, fullName := userSession(c)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"name":      "Авторизация",
		"id":        id,
		"full_name": fullName,
	})
}

func registration(c *gin.Context) {
	id, fullName := userSession(c)

	c.HTML(http.StatusOK, "registration.html", gin.H{
		"name":      "Регистрация",
		"id":        id,
		"full_name": fullName,
	})
}

func profile(c *gin.Context) {
	id, fullName := userSession(c)

	userActive := user.GetUserById(id)

	rool := userActive.Rool

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"name": fullName,
		"id":   id,
		"rool": rool,
	})
}

func profileUpdate(c *gin.Context) {
	id, fullName := userSession(c)

	user := user.GetUserById(id)

	first := user.FirstName
	last := user.LastName

	c.HTML(http.StatusOK, "update.html", gin.H{
		"name":  fullName,
		"id":    id,
		"first": first,
		"last":  last,
	})
}
