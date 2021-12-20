package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"myLibrary/pkg/handlers/book"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addBookProcess(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")
	url := c.PostForm("url")
	file := c.PostForm("file")

	bytesRead, err := ioutil.ReadFile(url + "\\" + file)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("pkg/books/text/"+file, bytesRead, 0755)
	if err != nil {
		log.Fatal(err)
	}

	book.AddBook(name, author, file)
	c.Redirect(http.StatusFound, "/book")
}

func getAllBooksProcess(c *gin.Context) {
	id, fullName := userSession(c)

	books := book.GetAllBooks()

	c.HTML(http.StatusOK, "books.html", gin.H{
		"name":      "Все книги",
		"id":        id,
		"full_name": fullName,
		"books":     books,
	})
}

func getNameBookProcess(c *gin.Context) {
	id, fullName := userSession(c)

	search := c.PostForm("search")

	books := book.GetAllNameBooks(search)

	c.HTML(http.StatusOK, "books.html", gin.H{
		"name":      "Все книги",
		"id":        id,
		"full_name": fullName,
		"books":     books,
	})
}

func getBookProcess(c *gin.Context) {
	id, fullName := userSession(c)
	idBook := c.Param("id")
	iduBook := convertUInt(idBook)

	book := book.GetBookById(iduBook)
	fileName := book.FileName

	log.Println(fileName)
	file, err := os.Open("./pkg/books/text/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""

	for scanner.Scan() {
		content += scanner.Text() + "<br>"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "book.html", gin.H{
		"name":      "Все книги",
		"id":        id,
		"full_name": fullName,
		"content":   content,
	})
}

func updateBookProcess(c *gin.Context) {
	var newMark float64
	idBook := c.Param("id")
	iduBook := convertUInt(idBook)
	userMark := c.PostForm("mark")

	userMarkNumber, err := strconv.ParseFloat(userMark, 64)
	if err != nil {
		log.Fatal(err)
	}

	bookInfo := book.GetBookById(iduBook)
	numberMark := float64(bookInfo.Number_Mark)
	mark := bookInfo.Mark

	markNumber, err := strconv.ParseFloat(mark, 64)
	if err != nil {
		log.Fatal(err)
	}

	newMark = ((numberMark * markNumber) + userMarkNumber) / (numberMark + 1)
	fmt.Println(newMark)
	strNewMark := fmt.Sprintf("%f", newMark)
	book.UpdateBook(iduBook, strNewMark)
}
