package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/navneet7716/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {

	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
}
