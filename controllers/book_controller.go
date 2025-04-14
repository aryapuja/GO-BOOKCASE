package controllers

import (
	"go-bookcase/models"
	"go-bookcase/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetAllBooks
// @Description Endpoint untuk menampilkan semua buku
// @Success 200 {array} models.Book "List buku"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/books [get]
func GetAllBooks(c *gin.Context) {
	books, err := services.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary GetBookInfo
// @Description Endpoint untuk menampilkan detail buku berdasarkan ID
// @Param id path int true "ID Buku"
// @Success 200 {object} models.Book "Detail Buku"
// @Failure 404 {object} map[string]interface{} "Buku tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/books/{id} [get]
func GetBookInfo(c *gin.Context) {
	bookID := c.Param("id")
	book, err := services.GetBookInfo(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary CreateBook
// @Description Endpoint untuk menambah buku baru
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Data Buku"
// @Success 200 {object} models.Book "Buku berhasil dibuat"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/books [post]
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

// @Summary UpdateBook
// @Description Endpoint untuk update buku berdasarkan ID
// @Param id path int true "ID Buku"
// @Param book body models.Book true "Update data buku"
// @Success 200 {object} models.Book "Data buku terupdate"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Buku tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/books/{id} [put]
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

// @Summary DeleteBook
// @Description Endpoint untuk menghapus buku berdasarkan ID
// @Param id path int true "ID Buku"
// @Success 200 {object} map[string]interface{} "Buku berhasil dihapus"
// @Failure 404 {object} map[string]interface{} "Buku tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	bookID := c.Param("id")
	err := services.DeleteBook(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
