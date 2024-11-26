package main

import (
	"log"
	"net/http"
)

func StartServer() {
	//Вывод пользователей из бд возможные параметры: page-страница данных;limit-кол-во юзеров на странице;name-фильтр по имени
	initDB() // Инициализация базы данных
	mux := http.NewServeMux()
	//обработка маршрутов для users(GET,POST(json))
	mux.HandleFunc("/users", UsersHandler)
	//обработка маршрутов для users/(GET,PUT(json)),DELETE+middleware+token
	//mux.HandleFunc("/users/", UserHandler)
	mux.Handle("/users/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			AuthMiddleware(http.HandlerFunc(UserHandler)).ServeHTTP(w, r)
		} else {
			UserHandler(w, r)
		}
	}))
	mux.Handle("/login", http.HandlerFunc(LoginHandler))
	handlerWithCORS := corsMiddleware(mux)
	log.Println("Server is running on http://localhost:8090")
	if err := http.ListenAndServe(":8090", handlerWithCORS); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

// Middleware для поддержки CORS
// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любых источников
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это предзапрос (OPTIONS), то просто отвечаем с 200 статусом
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
