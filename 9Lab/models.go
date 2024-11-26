package main

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Структура для входа
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Структура для ответа с токеном
type TokenResponse struct {
	Token string `json:"token"`
}
