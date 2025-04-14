package routes

import (
	"go-bookcase/controllers"
	"go-bookcase/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Route untuk login
	r.POST("/api/users/login", controllers.LoginUser)

	// Routes dengan middleware JWT
	authorized := r.Group("/")
	authorized.Use(middleware.JWTMiddleware()) // Apply JWT middleware for these routes

	// CRUD kategori
	authorized.GET("/api/categories", controllers.GetAllCategories)
	authorized.POST("/api/categories", controllers.CreateCategory)
	authorized.GET("/api/categories/:id", controllers.GetCategoryInfo)
	authorized.DELETE("/api/categories/:id", controllers.DeleteCategory)
	authorized.GET("/api/categories/:id/books", controllers.GetBooksByCategory)

	// CRUD buku
	authorized.GET("/api/books", controllers.GetAllBooks)
	authorized.POST("/api/books", controllers.CreateBook)
	authorized.GET("/api/books/:id", controllers.GetBookInfo)
	authorized.DELETE("/api/books/:id", controllers.DeleteBook)
}
