package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneet7716/go-bookstore/pkg/models"
	"github.com/navneet7716/go-bookstore/pkg/utils"
)

var newUser models.User

func GetUser(c *gin.Context) {

	newUser := models.GetAllUsers()
	c.IndentedJSON(http.StatusOK, newUser)

}

func CreateUser(c *gin.Context) {

	NewUser := &models.User{}
	utils.ParseBody(c.Request, NewUser)
	u := NewUser.CreateUser()
	c.IndentedJSON(http.StatusOK, u)

}
