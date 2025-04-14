package services

import (
	"database/sql"
	"go-bookcase/config"
	"go-bookcase/models"
)

// GetAllCategories retrieves all categories from the database
func GetAllCategories() ([]models.Category, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, name, created_at, created_by FROM categories`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// CreateCategory creates a new category in the database
func CreateCategory(category models.Category) (models.Category, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Category{}, err
	}
	defer db.Close()

	query := `INSERT INTO categories (name, created_by) VALUES ($1, $2) RETURNING id`
	var id int
	err = db.QueryRow(query, category.Name, category.CreatedBy).Scan(&id)
	if err != nil {
		return models.Category{}, err
	}

	category.ID = id
	return category, nil
}

// GetCategoryInfo retrieves category details by ID
func GetCategoryInfo(id string) (models.Category, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Category{}, err
	}
	defer db.Close()

	query := `SELECT id, name, created_at, created_by FROM categories WHERE id = $1`
	var category models.Category
	err = db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, err
		}
		return models.Category{}, err
	}

	return category, nil
}

// DeleteCategory deletes a category from the database
func DeleteCategory(id string) error {
	db, err := config.SetupDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM categories WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetBooksByCategory retrieves all books for a specific category
func GetBooksByCategory(categoryID string) ([]models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by 
			  FROM books WHERE category_id = $1`
	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}