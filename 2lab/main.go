/*
6)Написать функцию, которая принимает два целых числа и возвращает их среднее значение.
*/
package main

import (
	"fmt"
)

// 1)Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

// 2)Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".
func findSymbol(a float64) {
	if a == 0 {
		fmt.Printf("%f is Zero", a)
	} else if a > 0 {
		fmt.Printf("%f is Positive", a)
	} else if a < 0 {
		fmt.Printf("%f is Negative", a)
	}
}

// 4)Написать функцию, которая принимает строку и возвращает ее длину.
func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

// 5)Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

// 6)Написать функцию, которая принимает два целых числа и возвращает их среднее значение.
func average(a, b int) float64 {
	return float64(a+b) / 2.0
}
func main() {
	NumOne := 2
	NumTwo := 2.5
	NumThree := 3
	string := "Hello"
	isEven := isEven(NumOne)
	//1
	fmt.Printf("isEven(NumOne)=%t\n", isEven)
	//2
	findSymbol(NumTwo)
	//3)Написать программу, которая выводит все числа от 1 до 10 с помощью цикла for.
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//4
	fmt.Printf("In string %s %d symbols\n", string, strLen(string))
	//5
	rectangle := Rectangle{
		width:  10,
		height: 5,
	}
	fmt.Printf("Area of rectangle is %f\n", area(rectangle))
	//6
	fmt.Printf("Average of %d and %d is %f", NumOne, NumThree, average(NumOne, NumThree))
}
