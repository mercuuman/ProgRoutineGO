package main

//Обработчики маршрутов
import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Обработчики для /users(GET,POST)

// GetUsersHandler Получение пользователей из дб возможные параметры: page-страница данных;limit-кол-во юзеров на странице;name-фильтр по имени
// дефолтные параметры:1,10,""
// Пример: http://localhost:8080/users?page=2&limit=20&name=John
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	nameFilter := r.URL.Query().Get("name")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	users, err := getUsersFromDB(nameFilter, limit, page)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUserHandler принимает из POST json юзера и записывает его в бд
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	if err := validateUser(&user); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	newUser, err := createUserInDB(&user)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Обработчик для /users Разводит users на методы GET и POST
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("Получен запрос на GET /users")
		GetUsersHandler(w, r) // Получение списка пользователей
	} else if r.Method == http.MethodPost {
		log.Println("Получен запрос на POST /users")
		CreateUserHandler(w, r) // Создание нового пользователя
	} else {
		handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
	}
}

// Обработчики для /users/(GET,PUT,DELETE)

// GetUserHandler берет id и ищет по нему юзера в бд; метод GET
// Пример http://localhost:8080/users/1
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid user ID"), http.StatusBadRequest)
		return
	}

	user, err := getUserFromDB(id)
	if err != nil {
		handleError(w, err, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler принимает json юзера по id и обновляет данные в бд; метод PUT
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid user ID"), http.StatusBadRequest)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	if err := validateUser(&user); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	updatedUser, err := updateUserInDB(id, &user)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUserHandler Удаляет пользователя в бд по id; метод DELETE
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, fmt.Errorf("invalid user ID"), http.StatusBadRequest)
		return
	}

	err = deleteUserFromDB(id)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Обработчик для /users/ Разводит users/ на методы GET , PUT и DELETE
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetUserHandler(w, r) // Получение пользователя по ID
	} else if r.Method == http.MethodPut {
		UpdateUserHandler(w, r) // Обновление данных юзера
	} else if r.Method == http.MethodDelete {
		DeleteUserHandler(w, r) // Удаление пользователя
	} else {
		handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
	}
}

// Функционал 9 лабы
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	exists, err := findUser(&user)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	if !exists {
		handleError(w, fmt.Errorf("user not found"), http.StatusNotFound)
		return
	}
	// Генерация JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,                           // Дополнительные данные
		"exp":   time.Now().Add(1 * time.Hour).Unix(), // Срок действия токена
	})

	// Подписываем токен
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Возвращаем токен клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
