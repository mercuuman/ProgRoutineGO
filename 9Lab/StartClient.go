package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func StartClient() {

	fmt.Println("Клиент готов к работе. Введите команды:")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите 'login' для того чтобы залогиниться ")
		//маршруты /users
		fmt.Println("Введите 'list' для вывода списка юзеров ")
		fmt.Println("Введите 'IdU' для вывода юзера по id ")
		//маршруты /users/
		fmt.Println("Введите 'newU' для записи нового юзера ")
		fmt.Println("Введите 'UpdU' для обновления данных юзера ")
		fmt.Println("Введите 'DelU' для удаления юзера ")
		fmt.Println("Введите 'exit' для выхода")
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "login":
			login()
		case "list": //GET /users
			listUsers()
		case "IdU": //GET /users/{id}
			listUserById()
		case "newU": //POST /users
			newUser()
		case "UpdU": //PUT /users/
			updateUser()
		case "DelU": //PUT /users/
			deleteUser() //DELETE /users/
		case "exit": //
			fmt.Println("Завершение работы клиента...")
			return
		default:
			fmt.Println("Неизвестная команда. Используйте 'list' или 'exit'.")
		}
	}
}

func listUsers() {
	client := &http.Client{}
	fmt.Println("Запрос списка пользователей...")

	req, err := http.NewRequest("GET", "http://localhost:8090/users", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Сервер вернул ошибку: %d %s\n", resp.StatusCode, string(body))
		return
	}

	// Распарсим JSON-ответ в массив пользователей
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}

	// Выводим список пользователей в консоль
	fmt.Println("Список пользователей:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
func newUser() {
	var email, password, name string

	fmt.Println("Введите ваш email:")
	fmt.Scanln(&email)

	fmt.Println("Введите ваш пароль:")
	fmt.Scanln(&password)

	fmt.Println("Введите ваше имя:")
	fmt.Scanln(&name)
	user := User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Ошибка при преобразовании данных в JSON:", err)
	}
	resp, err := http.Post("http://localhost:8090/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Fatal("Ошибка при отправке запроса:", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении ответа:", err)
	}

	// Вывод ответа сервера
	fmt.Printf("Ответ сервера: %s\n", body)
}
func listUserById() {
	var id string

	fmt.Println("Введите id :")
	fmt.Scanln(&id)
	url := fmt.Sprintf("http://localhost:8090/users/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка при отправке запроса:", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении ответа:", err)
	}

	// Вывод ответа сервера
	fmt.Printf("Ответ сервера: %s\n", body)
}
func updateUser() {
	var email, password, name, id string
	fmt.Println("Введите id юзера:")
	fmt.Scanln(&id)
	// Запрос данных пользователя
	fmt.Println("Введите ваш email:")
	fmt.Scanln(&email)

	fmt.Println("Введите ваш пароль:")
	fmt.Scanln(&password)

	fmt.Println("Введите ваше имя:")
	fmt.Scanln(&name)
	url := fmt.Sprintf("http://localhost:8090/users/%s", id)
	// Создание структуры User с введенными данными
	user := User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	// Преобразование данных в JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Ошибка при преобразовании данных в JSON:", err)
	}

	// Создание PUT-запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(userJSON))
	if err != nil {
		log.Fatal("Ошибка при создании запроса:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Ошибка при отправке запроса:", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении ответа:", err)
	}

	// Вывод ответа сервера
	fmt.Printf("Ответ сервера: %s\n", body)
}
func deleteUser() {
	var id string

	token, err := getTokenFromFile()
	if err != nil {
		log.Fatal("Ошибка при чтении токена:", err)
	}

	// Запрос id пользователя
	fmt.Println("Введите id :")
	fmt.Scanln(&id)

	// Формирование URL для удаления пользователя
	url := fmt.Sprintf("http://localhost:8090/users/%s", id)

	// Создание DELETE-запроса
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Ошибка при создании запроса:", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	// Создание HTTP-клиента
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Ошибка при отправке запроса:", err)
	}
	defer resp.Body.Close()

	// Чтение ответа от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении ответа:", err)
	}

	// Вывод ответа сервера
	fmt.Printf("Ответ сервера: %s\n", body)
}
func login() {
	var email, password, name string

	fmt.Println("Введите вашюзернейм:")
	fmt.Scanln(&name)

	fmt.Println("Введите ваш пароль:")
	fmt.Scanln(&password)

	// Формируем JSON-запрос
	user := User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Ошибка при преобразовании данных в JSON:", err)
	}

	// Отправляем POST-запрос
	resp, err := http.Post("http://localhost:8090/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Ошибка при отправке запроса:", err)
	}
	defer resp.Body.Close()

	// Читаем ответ сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при чтении ответа:", err)
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка авторизации: %s\n", string(body))
		return
	}

	// Извлекаем токен из ответа
	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}

	token, ok := response["token"]
	if !ok {
		fmt.Println("Ответ не содержит токен.")
		return
	}

	// Сохраняем токен
	err = saveTokenToFile(token)
	if err != nil {
		log.Println("Не удалось сохранить токен:", err)
	} else {
		fmt.Println("Токен сохранён в session.txt")
	}

	fmt.Println("Авторизация успешна!")
}
func saveTokenToFile(token string) error {
	return ioutil.WriteFile("session.txt", []byte(token), 0644)
}
func getTokenFromFile() (string, error) {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
