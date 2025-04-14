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

// CreateBook - Endpoint untuk menambah buku
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := services.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdBook)
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
