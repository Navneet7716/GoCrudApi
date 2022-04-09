package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/navneet7716/go-bookstore/pkg/controllers"
)

var RegisterUserRoutes = func(router *gin.Engine) {

	router.GET("/user/", controllers.GetUser)
	router.POST("/user/", controllers.CreateUser)
}
