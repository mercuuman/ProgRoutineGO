/*
1.Создание и запуск горутин:
•	Напишите программу, которая параллельно выполняет три функции (например, расчёт факториала, генерация случайных чисел и вычисление суммы числового ряда).
•	Каждая функция должна выполняться в своей горутине.
•	Добавьте использование time.Sleep() для имитации задержек и продемонстрируйте параллельное выполнение.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Factorial(a int, wg *sync.WaitGroup) int {
	defer wg.Done()
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
		fmt.Printf("Factorial of %d is %d \n", i, result)
	}
	time.Sleep(150 * time.Millisecond)
	return result
}
func RanNum(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= a; i++ {
		result := rand.Intn(10)
		fmt.Printf(" %d random number is %d \n", i, result)
	}
	time.Sleep(100 * time.Millisecond)
}
func NumSum(a int, wg *sync.WaitGroup) int {
	defer wg.Done()
	result := 0
	for i := 1; i <= a; i++ {
		result += i
		fmt.Printf(" sum of 1-%d  %d \n", i, result)
	}
	time.Sleep(200 * time.Millisecond)
	return result
}

func main() {
	var wg sync.WaitGroup
	a := 5
	wg.Add(1)
	go Factorial(a, &wg)
	wg.Add(1)
	go NumSum(a, &wg)
	wg.Add(1)
	go RanNum(a, &wg)
	wg.Wait()
}
