// auth.go
package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("your_secret_key")

// Генерация JWT
func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Проверка JWT
func validateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неизвестный метод подписи")
		}
		return jwtSecret, nil
	})

	return token, err
}
