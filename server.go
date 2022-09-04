package main

import (
	"log"

	"github.com/Dominggus11/MyPROject/book"
	"github.com/Dominggus11/MyPROject/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=rdam password=programming dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Error")
	}

	//Auto Migrate pada database
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	//membuat router
	router := gin.Default()

	//membuat grup router
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)
	// router.Run(":8888") make port 8888
	router.Run() //make port default 8080
}
