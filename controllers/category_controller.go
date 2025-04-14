package controllers

import (
	"go-bookcase/models"
	"go-bookcase/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetAllCategories
// @Description Endpoint untuk menampilkan semua kategori
// @Success 200 {array} models.Category "List Kategori"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories [get]
func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary GetCategoryInfo
// @Description Endpoint untuk menampilkan detail kategori berdasarkan ID
// @Param id path int true "ID Kategori"
// @Success 200 {object} models.Category "Info Kategori"
// @Failure 404 {object} map[string]interface{} "Kategori tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id} [get]
func GetCategoryInfo(c *gin.Context) {
	categoryID := c.Param("id")
	category, err := services.GetCategoryInfo(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Summary GetBooksByCategory
// @Description Endpoint untuk menampilkan buku berdasarkan kategori
// @Param id path int true "ID Kategori"
// @Success 200 {array} models.Book "List buku sesuai kategori"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id}/books [get]
func GetBooksByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	books, err := services.GetBooksByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary CreateCategory
// @Description Endpoint untuk menambah kategori baru
// @Accept  json
// @Produce  json
// @Param category body models.Category true "Data Kategoti"
// @Success 200 {object} models.Category "Kategori berhasil dibuat"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories [post]
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

// @Summary UpdateCategory
// @Description Endpoint untuk update kategori berdasarkan ID
// @Param id path int true "ID Kategori"
// @Param category body models.Category true "Updated data kategori"
// @Success 200 {object} models.Category "Kategori terupdate"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Router /api/categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := services.UpdateCategory(categoryID, category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Update Failed: ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// @Summary DeleteCategory
// @Description Endpoint untuk menghapus kategori berdasarkan ID
// @Param id path int true "ID Kategori"
// @Success 200 {object} map[string]interface{} "Kategori dihapus"
// @Failure 404 {object} map[string]interface{} "Kategori tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")
	err := services.DeleteCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delete Failed: ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
