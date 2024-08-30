package handlers

import (
	"net/http"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
)

// CreateBook handles creating a new book
func CreateBook(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid book data",
		})
	}

	result := db.Create(book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create book",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Book created successfully",
		"book_id": book.ID,
	})
}

// GetBook handles getting a book by ID
func GetBook(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid book ID",
		})
	}

	book := new(models.Book)
	result := db.First(book, bookID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Book not found",
		})
	}

	return c.JSON(http.StatusOK, book)
}

// GetAllBooks handles getting all books
func GetAllBooks(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var books []models.Book
	result := db.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve books",
		})
	}

	return c.JSON(http.StatusOK, books)
}