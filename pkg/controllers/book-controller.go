package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/navneet7716/go-bookstore/pkg/models"
	"github.com/navneet7716/go-bookstore/pkg/utils"
)

var newBook models.Book

func GetBook(c *gin.Context) {

	newBooks := models.GetAllBooks()
	c.IndentedJSON(http.StatusOK, newBooks)

}

func GetBookById(c *gin.Context) {

	ID := c.Param("id")

	id, err := strconv.ParseInt(ID, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Some error while getting the data."})
	}

	book, _ := models.GetBookById(id)

	if book.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book with this id exists."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	utils.ParseBody(c.Request, CreateBook)
	b := CreateBook.CreateBook()
	c.IndentedJSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	ID := c.Param("id")

	id, err := strconv.ParseInt(ID, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Some error while getting the data."})
	}

	models.DeleteBook(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data Deleted."})
}

func UpdateBook(c *gin.Context) {
	var updatedBook = &models.Book{}
	utils.ParseBody(c.Request, updatedBook)
	ID := c.Param("id")

	id, err := strconv.ParseInt(ID, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Some error while getting the data."})
	}

	bookDetails, db := models.GetBookById(id)

	if bookDetails.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book with this id exists."})
		return
	}

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db.Save(&bookDetails)
	c.IndentedJSON(http.StatusCreated, updatedBook)

}
