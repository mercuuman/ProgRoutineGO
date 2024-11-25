package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

// Инициализация бд
func initDB() {
	var err error
	connStr := "user=youser password=641 dbname=lab8 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to the database!")
}

// Создание пользователя по структуре user
func createUserInDB(user *User) (User, error) {
	var newUser User
	query := "INSERT INTO users(name, email,password) VALUES($1, $2, $3) RETURNING id"
	err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&newUser.ID)
	if err != nil {
		return newUser, err
	}
	newUser.Name = user.Name
	newUser.Email = user.Email
	return newUser, nil
}

// Вывод юзеров с дб
// nameFilter - строка для фильтра по имени(можно передать пустую | limit - юзеры на 1 странице | page - номер страницы
func getUsersFromDB(nameFilter string, limit, page int) ([]User, error) {
	offset := (page - 1) * limit
	query := "SELECT id, name, email FROM users WHERE name ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3"
	rows, err := db.Query(query, nameFilter, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Вывод юзера по id
func getUserFromDB(id int) (User, error) {
	var user User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Обновление данных юзера. Передаем юзера возвращаем его такого же если все прошло окей
func updateUserInDB(id int, user *User) (User, error) {
	var updatedUser User
	query := "UPDATE users SET name = $1, email = $2, password=$3 WHERE id=$4 RETURNING id, name, email,password"
	err := db.QueryRow(query, user.Name, user.Email, user.Password, id).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.Password)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

// Удаление юзера из дб по id
func deleteUserFromDB(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}
