package services

import (
	"database/sql"
	"errors"
	"go-bookcase/config"
	"go-bookcase/models"
)

// retrieves all book
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
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Retrieves book details by ID
func GetBookInfo(id string) (models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by FROM books WHERE id = $1`
	var book models.Book
	err = db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, err
		}
		return models.Book{}, err
	}

	return book, nil
}

// Creates a new book
func CreateBook(book models.Book) (models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	query := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at`

	err = db.QueryRow(query, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy).
		Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, errors.New("failed to insert book")
		}
		return models.Book{}, err
	}

	return book, nil
}

// Update book by ID
func UpdateBook(id string, book models.Book) (models.Book, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return models.Book{}, err
	}
	defer db.Close()

	var existingBook models.Book
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&existingBook.ID, &existingBook.Title, &existingBook.Description, &existingBook.ImageUrl, &existingBook.ReleaseYear, &existingBook.Price, &existingBook.TotalPage, &existingBook.Thickness, &existingBook.CategoryID)
	if err == sql.ErrNoRows {
		return models.Book{}, errors.New("book not found")
	} else if err != nil {
		return models.Book{}, err
	}

	updateQuery := `UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, modified_at = current_timestamp WHERE id = $9 RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, modified_at`
	err = db.QueryRow(updateQuery, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, id).
		Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.ModifiedAt)

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

// Deletes a book by ID
func DeleteBook(id string) error {
	db, err := config.SetupDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	var existingBook models.Book
	query := `SELECT id FROM books WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&existingBook.ID)
	if err == sql.ErrNoRows {
		return errors.New("book not found")
	} else if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM books WHERE id = $1`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
