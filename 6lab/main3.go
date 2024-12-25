/*
3.	Применение select для управления каналами:
•	Создайте две горутины, одна из которых будет генерировать случайные числа, а другая — отправлять сообщения об их чётности/нечётности.
•	Используйте кон
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RandNum(a int, chanel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= a; i++ {
		result := rand.Intn(10)
		fmt.Printf("[RandNum] Generated: %d\n", result)
		chanel <- result // Отправляем число в канал
		time.Sleep(500 * time.Millisecond)
	}
	close(chanel) // Закрываем канал после отправки всех чисел
}

func EvenOdd(chanelInt chan int, chanelStr chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range chanelInt { // Читаем числа из канала
		if num%2 == 0 {
			chanelStr <- fmt.Sprintf("Number %d is even", num)
		} else {
			chanelStr <- fmt.Sprintf("Number %d is odd", num)
		}
	}
	close(chanelStr) // Закрываем канал после отправки всех сообщений
}

func Selecter(chanelInt chan int, chanelStr chan string) {
	for {
		select {
		case valueInt, ok := <-chanelInt:
			if ok {
				fmt.Printf("[Selecter] Received number from RandNum: %d\n", valueInt)
			} else {
				chanelInt = nil // Ставим канал в nil, чтобы предотвратить дальнейшие операции
			}
		case strOut, ok := <-chanelStr:
			if ok {
				fmt.Printf("[Selecter] Received message from EvenOdd: %s\n", strOut)
			} else {
				chanelStr = nil // Ставим канал в nil, чтобы предотвратить дальнейшие операции
			}
		}

		// Проверяем, закрыты ли оба канала
		if chanelInt == nil && chanelStr == nil {
			fmt.Println("[Selecter] Both channels are closed. Exiting.")
			break
		}
	}
}

func main() {
	a := 10
	var wg sync.WaitGroup
	chanelInt := make(chan int)
	chanelStr := make(chan string)

	wg.Add(1)
	go RandNum(a, chanelInt, &wg)
	wg.Add(1)
	go EvenOdd(chanelInt, chanelStr, &wg)

	go Selecter(chanelInt, chanelStr)

	wg.Wait()
}
