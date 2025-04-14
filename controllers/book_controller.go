package controllers

import (
	"go-bookcase/models"
	"go-bookcase/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBooks - Endpoint untuk menampilkan semua buku
func GetAllBooks(c *gin.Context) {
	books, err := services.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBookInfo - Endpoint untuk menampilkan detail buku
func GetBookInfo(c *gin.Context) {
	bookID := c.Param("id")
	book, err := services.GetBookInfo(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook - Endpoint untuk menambah buku
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	createdBook, err := services.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}

// UpdateBook - Endpoint untuk update buku
func UpdateBook(c *gin.Context) {
	bookID := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	updatedBook, err := services.UpdateBook(bookID, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid data / Book Not Found"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook - Endpoint untuk menghapus buku
func DeleteBook(c *gin.Context) {
	bookID := c.Param("id")
	err := services.DeleteBook(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
