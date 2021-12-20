package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func funcHandler() {
	store := cookie.NewStore([]byte("secret"))
	router := gin.Default()
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("ui/html/*")

	router.GET("/", home)
	router.GET("/login", login)
	router.POST("/loginuser", loginuser)
	router.GET("/registration", registration)
	router.GET("/profile/update/:id", profileUpdate)
	router.GET("/profile/exit/:id", profileExit)
	router.POST("/profile/addbook", addBookProcess)
	router.POST("/search", getNameBookProcess)

	user := router.Group("/user")
	{
		user.POST("/", addUserProcess)
		user.GET("/", getAllUsersProcess)
		user.GET("/:id", profile)
		user.PUT("/name/:id", updateUserNameProcess)
		user.PUT("/password/:id", updateUserPasswordProcess)
		user.DELETE("/:id", deleteUserProcess)
	}

	book := router.Group("/book")
	{
		book.POST("", addBookProcess)
		book.GET("", getAllBooksProcess)
		book.GET("/:id", getBookProcess)
		book.PUT("/:id", updateBookProcess)
	}

	router.Run(":8080")
}
