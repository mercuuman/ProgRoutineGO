/*
2.	Использование каналов для передачи данных:
•	Реализуйте приложение, в котором одна горутина генерирует последовательность чисел (например, первые 10 чисел Фибоначчи) и отправляет их в канал.
•	Другая горутина должна считывать данные из канала и выводить их на экран.
•	Добавьте блокировку чтения из канала с помощью close() и объясните её роль.
*/
package main

import (
	"fmt"
	"sync"
)

func Fibonacci(a int, chanel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for i := 0; i < a; i++ {
		chanel <- x
		fmt.Printf(" %d  number is %d, send do chanel\n", i, x)
		x, y = y, x+y
	}
	close(chanel)
}

func PrintNum(chanel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range chanel {
		fmt.Printf("Received random number: %d\n", num)
	}
}
func main() {
	a := 10
	var wg sync.WaitGroup
	chanel := make(chan int)
	wg.Add(1)
	go Fibonacci(a, chanel, &wg)
	wg.Add(1)
	go PrintNum(chanel, &wg)
	wg.Wait()

}
