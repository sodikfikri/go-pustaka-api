package main

import (
	"pustaka-api/book"
	dbconnection "pustaka-api/dbConnection"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db := dbconnection.DBInit()

	booRepsitory := book.NewRepository(db)
	bookService := book.NewService(booRepsitory)
	bookHandler := handler.NewBookHandler((bookService))

	route := gin.Default()
	// v1 := route.G

	// route.GET("/", bookHandler.RouteHandler)
	// route.GET("/book/:id", bookHandler.BookHandler)
	// route.GET("/query", bookHandler.QueryHandler)
	route.GET("/books/:id", bookHandler.GetByID)
	route.GET("/books", bookHandler.GetAll)
	route.POST("/books", bookHandler.Create)
	route.PUT("/books/:id", bookHandler.Update)

	route.Run(":8888")
}
