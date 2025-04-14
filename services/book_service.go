package services

import (
	"database/sql"
	"go-bookcase/config"
	"go-bookcase/models"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks() ([]models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by FROM books`
	rows, err := db.Query(query)
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

// CreateBook creates a new book in the database
func CreateBook(book models.Book) (models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	query := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	var id int
	err = db.QueryRow(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy).Scan(&id)
	if err != nil {
		return models.Book{}, err
	}

	book.ID = id
	return book, nil
}

// GetBookInfo retrieves book details by ID
func GetBookInfo(id string) (models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by FROM books WHERE id = $1`
	var book models.Book
	err = db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, err
		}
		return models.Book{}, err
	}

	return book, nil
}

// DeleteBook deletes a book from the database
func DeleteBook(id string) error {
	db, err := config.SetupDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM books WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
