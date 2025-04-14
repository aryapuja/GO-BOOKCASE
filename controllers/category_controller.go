package controllers

import (
	"go-bookcase/models"
	"go-bookcase/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCategories - Endpoint untuk menampilkan semua kategori
func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetBooksByCategory - Endpoint untuk menampilkan buku berdasarkan kategori
func GetBooksByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	books, err := services.GetBooksByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// CreateCategory - Endpoint untuk menambah kategori
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdCategory)
}

// GetCategoryInfo - Endpoint untuk menampilkan detail kategori
func GetCategoryInfo(c *gin.Context) {
	categoryID := c.Param("id")
	category, err := services.GetCategoryInfo(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory - Endpoint untuk update kategori
func UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := services.UpdateCategory(categoryID, category)
	if err != nil {
		// Mengubah pesan error menjadi lebih jelas
		c.JSON(http.StatusNotFound, gin.H{"error": "Update Failed: ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory - Endpoint untuk menghapus kategori
func DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")
	err := services.DeleteCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delete Failed: ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
