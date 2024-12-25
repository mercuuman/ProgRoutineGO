/*
6.	Создание пула воркеров:
•	Реализуйте пул воркеров, обрабатывающих задачи (например, чтение строк из файла и их реверсирование).
•	Количество воркеров задаётся пользователем.
•	Распределение задач и сбор результатов осуществляется через каналы.
•	Выведите результаты работы воркеров в итоговый файл или в консоль.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Воркеры реверсируют строку и возвращают результат через канал
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Реверсирование строки
		reversed := reverseString(job)
		fmt.Printf("Worker %d processed string: %s -> %s\n", id, job, reversed)
		// Отправляем результат
		results <- reversed
	}
}

// Функция для реверсирования строки
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Открываем файл для чтения строк
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Чтение строк из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Пользователь задает количество воркеров
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	jobs := make(chan string, len(lines))
	results := make(chan string, len(lines))

	// Синхронизация воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Отправляем задачи (строки) в канал jobs
	for _, line := range lines {
		jobs <- line
	}
	close(jobs)
	wg.Wait()
	close(results)

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	// Считываем результаты из канала и выводим в файл
	for result := range results {
		_, err := outputFile.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}

	fmt.Println("Обработка завершена, результаты сохранены в файл output.txt")
}
