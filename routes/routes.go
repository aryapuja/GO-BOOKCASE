package routes

import (
	"go-bookcase/controllers"
	"go-bookcase/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Route login
	r.POST("/api/users/login", controllers.LoginUser)

	// Routes middleware JWT
	authorized := r.Group("/")
	authorized.Use(middleware.JWTMiddleware()) // Apply JWT middleware

	// CRUD kategori
	authorized.GET("/api/categories", controllers.GetAllCategories)
	authorized.GET("/api/categories/:id", controllers.GetCategoryInfo)
	authorized.GET("/api/categories/:id/books", controllers.GetBooksByCategory)
	authorized.POST("/api/categories", controllers.CreateCategory)
	authorized.PUT("/api/categories/:id", controllers.UpdateCategory)
	authorized.DELETE("/api/categories/:id", controllers.DeleteCategory)

	// CRUD buku
	authorized.GET("/api/books", controllers.GetAllBooks)
	authorized.GET("/api/books/:id", controllers.GetBookInfo)
	authorized.POST("/api/books", controllers.CreateBook)
	authorized.PUT("/api/books/:id", controllers.UpdateBook)
	authorized.DELETE("/api/books/:id", controllers.DeleteBook)
}
