package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
)

// Дополнительные функции для обработчиков
// Обработчик ошибок передаем ошибку и статус для ошибки
func handleError(w http.ResponseWriter, err error, status int) {
	log.Println("Error:", err) // Логирование ошибки
	w.WriteHeader(status)
	response := map[string]string{"error": err.Error()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Проверка на заполненность имейла и имени у юзера
func validateUser(user *User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

// Проверка правильности формата email
func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Функция хеширования пароля
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
