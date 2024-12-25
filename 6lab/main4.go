/*4.	Синхронизация с помощью мьютексов:
•	Реализуйте программу, в которой несколько горутин увеличивают общую переменную-счётчик.
•	Используйте мьютексы (sync.Mutex) для предотвращения гонки данных.
•	Включите и выключите мьютексы, чтобы увидеть разницу в работе программы.*/

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex // Мьютекс
)

func Increment() {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func incrementNoMutex() {
	defer wg.Done()
	counter++
}

func main() {
	// Пример с мьютексами
	counter = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go Increment() // Запускаем 100 горутин с мьютексами
	}
	wg.Wait()
	fmt.Println("Final Counter with Mutex:", counter)

	// Пример без мьютексов
	counter = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go incrementNoMutex() // Запускаем 100 горутин без мьютексов
	}
	wg.Wait()
	fmt.Println("Final Counter without Mutex:", counter)
}
