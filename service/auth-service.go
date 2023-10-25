package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Fungsi untuk membuat token JWT
func createToken(userID int) (string, error) {
	// Buat klaim
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Waktu kedaluwarsa token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
