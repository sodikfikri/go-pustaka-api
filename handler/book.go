package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// func (handler *bookHandler) RouteHandler(res *gin.Context) {
// 	res.JSON(http.StatusOK, gin.H{
// 		"name": "Sodik Fikri",
// 		"bio":  "Software Engineer",
// 	})
// }

// func (handler *bookHandler) BookHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	c.JSON(http.StatusOK, gin.H{"id": id})
// }

// func (handler *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")

// 	c.JSON(http.StatusOK, gin.H{"title": title})
// }

func (handler *bookHandler) GetAll(c *gin.Context) {
	book, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksRespose []responses.BookResponse

	for _, b := range book {

		bookRespose := convertToBookResponse(b)

		booksRespose = append(booksRespose, bookRespose)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksRespose,
	})
}

func (handler *bookHandler) GetByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	data, err := handler.bookService.FindByID(id)

	var booksRespose []responses.BookResponse

	resp := convertToBookResponse(data)

	booksRespose = append(booksRespose, resp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksRespose[0],
	})
}

func (handler *bookHandler) Create(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := handler.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (handler *bookHandler) Update(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := handler.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func convertToBookResponse(obj book.Book) responses.BookResponse {
	return responses.BookResponse{
		ID:          obj.ID,
		Title:       obj.Title,
		Price:       json.Number(strconv.Itoa(obj.Price)),
		Description: obj.Description,
		Rating:      obj.Rating,
		Discount:    obj.Discount,
	}
}
