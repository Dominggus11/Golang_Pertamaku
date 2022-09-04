package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dominggus11/MyPROject/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Roy Doms Andornov Malau",
		"Bio":  "Student Of Atmajaya University",
	})
}

func (handler *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
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

func (handler *bookHandler) GetBooks(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var BooksResponse []book.BooksResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		BooksResponse = append(BooksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": BooksResponse,
	})
}

func (handler *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := handler.bookService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BooksResponse {
	return book.BooksResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}

func (handler *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
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

func (handler *bookHandler) DeleteBook(c *gin.Context) {
	var bookRequest book.BookRequest
	// err := c.ShouldBindJSON(&bookRequest)
	// if err != nil {
	// 	errorMessages := []string{}
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
	// 		errorMessages = append(errorMessages, errorMessage)
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"errors": errorMessages,
	// 	})
	// 	return
	// }
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	_, err := handler.bookService.Delete(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "ID Tidak Ditemukan",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Data Berhasil Di Hapus",
	})

}
