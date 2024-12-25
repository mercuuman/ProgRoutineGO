/*
5.	Разработка многопоточного калькулятора:
•	Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /).
•	Используйте каналы для отправки запросов и возврата результатов.
•	Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.
http://localhost:8080/calculate?operation=-&num1=8&num2=5
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func calculate(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("operation")
	num1Str := r.URL.Query().Get("num1")
	num2Str := r.URL.Query().Get("num2")

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Ошибка: некорректные числа", http.StatusBadRequest)
		return
	}

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			http.Error(w, "Ошибка: деление на ноль", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	default:
		http.Error(w, "Ошибка: некорректная операция", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Результат: %.2f", result)
}

func main() {
	http.HandleFunc("/calculate", calculate)
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
