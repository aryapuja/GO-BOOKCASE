package services

import (
	"database/sql"
	"errors"
	"go-bookcase/config"
	"go-bookcase/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser checks: returns a JWT token
func AuthenticateUser(username, password string) (string, error) {
	db, err := config.SetupDatabase()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var user models.User
	query := `SELECT id, username, password FROM users WHERE username=$1`
	err = db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("INVALID PASSWORD)")
	}

	// Generate JWT Token if credentials valid
	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT generates a JWT token for a user
func generateJWT(user models.User) (string, error) {
	// Set claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Get the secret key from the .env
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("SECRET_KEY is not set in the environment variables")
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
