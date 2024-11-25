
### CLIENT
package main

import (
"bytes"
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
)

type User2 struct {
Name     string `json:"name"`
Email    string `json:"email"`
Password string `json:"password"`
}

func main() {
for {
// Меню
fmt.Println("Меню:")
fmt.Println("1. Добавить пользователя")
fmt.Println("2. Получить список пользователей")
fmt.Println("3. Получить информацию о пользователе")
fmt.Println("4. Обновить информацию о пользователе")
fmt.Println("5. Удалить пользователя")
fmt.Println("6. Выйти")
fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Добавить пользователя
			addUser()
		case 2:
			// Получить список пользователей
			getUsers()
		case 3:
			// Получить информацию о пользователе
			getUserInfo()
		case 4:
			// Обновить информацию о пользователе
			updateUser()
		case 5:
			// Удалить пользователя
			deleteUser()
		case 6:
			// Выход
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}

func addUser() {
var name, email, password string

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите email: ")
	fmt.Scan(&email)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)

	user := User2{Name: name, Email: email, Password: password}

	// Преобразование структуры в JSON
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Ошибка при маршалинге данных: %v", err)
	}

	// Отправка POST-запроса
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка при регистрации. Код ответа: %d", resp.StatusCode)
	} else {
		fmt.Println("Пользователь успешно зарегистрирован!")
	}
}

func getUsers() {
// Отправка GET-запроса для получения списка пользователей
resp, err := http.Get("http://localhost:8080/users")
if err != nil {
log.Fatalf("Ошибка при отправке запроса: %v", err)
}
defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", string(body))
}

func getUserInfo() {
var userId string
fmt.Print("Введите ID пользователя: ")
fmt.Scan(&userId)

	// Отправка GET-запроса для получения информации о пользователе
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s", userId))
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", string(body))
}

func updateUser() {
var userId, name, email, password string
fmt.Print("Введите ID пользователя для обновления: ")
fmt.Scan(&userId)
fmt.Print("Введите новое имя: ")
fmt.Scan(&name)
fmt.Print("Введите новый email: ")
fmt.Scan(&email)
fmt.Print("Введите новый пароль: ")
fmt.Scan(&password)

	// Создание структуры с новыми данными
	user := User2{Name: name, Email: email, Password: password}

	// Преобразование структуры в JSON
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Ошибка при маршалинге данных: %v", err)
	}

	// Отправка PUT-запроса для обновления информации о пользователе
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/users/%s", userId), bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", string(body))
}

func deleteUser() {
var userId string
fmt.Print("Введите ID пользователя для удаления: ")
fmt.Scan(&userId)

	// Отправка DELETE-запроса для удаления пользователя
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/users/%s", userId), nil)
	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", string(body))
}



### HANDLERS
package main

import (
"database/sql"
"encoding/json"
"fmt"
"golang.org/x/crypto/bcrypt"
"log"
"net/http"
"regexp"
"strconv"
)

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
func handleError(w http.ResponseWriter, err error, status int) {
log.Println("Error:", err) // Логирование ошибки
w.WriteHeader(status)
response := map[string]string{"error": err.Error()}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}

func validateUser(user *User) error {
if user.Name == "" {
return fmt.Errorf("name is required")
}
if user.Email == "" {
return fmt.Errorf("email is required")
}
return nil
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
return
}

	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// обработка данных и сохранение
	fmt.Fprintf(w, "User %s successfully registered with email %s", user.Name, user.Email)
}

// Функция для проверки правильности формата email
func isValidEmail(email string) bool {
// Регулярное выражение для проверки email
re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
return re.MatchString(email)
}

// Функция хеширования пароля
func hashPassword(password string) (string, error) {
// Хеширование с использованием bcrypt или другой библиотеки
// Пример с bcrypt (не забудьте подключить пакеты для bcrypt)
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
return "", err
}
return string(hashedPassword), nil
}

// Функция получения пользователя по email из базы данных
func getUserByEmail(email string) (User, error) {
var user User
query := "SELECT id, name, email FROM users WHERE email = $1"
err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email)
if err != nil {
if err == sql.ErrNoRows {
return user, fmt.Errorf("user not found")
}
return user, err
}
return user, nil
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
return
}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверьте пользователя в базе данных
	user, err := getUserByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Сравните пароли
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Генерация JWT токена
	token, err := generateJWT(user.ID)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


### MAIN
``package main

import (
"fmt"
"log"
"net/http"
"strconv"
)

func main() {

	initDB() // Инициализация базы данных

	mux := http.NewServeMux()

	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/users", usersHandler) // Обработка запросов к /users
	mux.HandleFunc("/users/", userHandler) // Обработка запросов к /users/{id}

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Обработчик для /users
func usersHandler(w http.ResponseWriter, r *http.Request) {
if r.Method == http.MethodGet {
GetUsersHandler(w, r) // Получение списка пользователей
} else if r.Method == http.MethodPost {
CreateUserHandler(w, r) // Создание нового пользователя
} else {
handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
}
}

// Обработчик для /users/{id}
func userHandler(w http.ResponseWriter, r *http.Request) {
idStr := r.URL.Path[len("/users/"):]
_, err := strconv.Atoi(idStr)
if err != nil {
handleError(w, fmt.Errorf("invalid user ID"), http.StatusBadRequest)
return
}

	switch r.Method {
	case http.MethodGet:
		GetUserHandler(w, r)
	case http.MethodPut:
		UpdateUserHandler(w, r)
	case http.MethodDelete:
		DeleteUserHandler(w, r)
	default:
		handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
	}
}``


### Auth
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



### DATABASE DONE

