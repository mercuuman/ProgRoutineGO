package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Флаг для выбора запуска
	mode := flag.String("mode", "", "Режим запуска: 'server' или 'client'")
	flag.Parse()

	if *mode == "server" {
		// Запуск сервера
		fmt.Println("Запускается сервер...")
		StartServer()
	} else if *mode == "client" {
		//Запуск клиента
		fmt.Println("Запускается клиент...")
		StartClient()
	} else {
		// Ошибка, если флаг не указан
		fmt.Println("Укажите режим запуска с помощью флага --mode (server или client).")
		os.Exit(1)
	}
}
