package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneet7716/go-bookstore/pkg/routes"
)

func main() {

	r := gin.Default()
	routes.RegisterBookStoreRoutes(r)
	routes.RegisterUserRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
