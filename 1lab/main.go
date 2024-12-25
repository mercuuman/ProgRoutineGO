package main

import (
	"fmt"
	"time"
)

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func sumAndDiff(a, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Функция для вычисления среднего значения трёх чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}
func main() {
	//1)Написать программу, которая выводит текущее время и дату.
	currentTime := time.Now()
	fmt.Println("Current time is equal:", currentTime.String())
	//Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
	var intVar int = 2
	var floatVar float64 = 6.5
	var stringVar string = "Word"
	var boolVar bool = true
	fmt.Println("Int variable:", intVar)
	fmt.Println("float variable:", floatVar)
	fmt.Println("string variable:", stringVar)
	fmt.Println("boolean variable:", boolVar)
	//Использовать краткую форму объявления переменных для создания и вывода переменных.+++++
	intVar2 := 4
	floatVar2 := 4.5
	stringVar2 := "Another Word"
	boolVar2 := false
	fmt.Println("Another variables!")
	fmt.Println(intVar2)
	fmt.Println(floatVar2)
	fmt.Println(stringVar2)
	fmt.Println(boolVar2)

	//Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
	a := 10
	b := 5
	fmt.Printf("%d + %d = %d\n\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	//Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
	sum, diff := sumAndDiff(floatVar, floatVar2)
	fmt.Printf("Sum is:%f , Diff is:%f\n", sum, diff)
	//Написать программу, которая вычисляет среднее значение трех чисел.
	floatVar3 := 4.7
	avg := average(floatVar, floatVar2, floatVar3)
	fmt.Printf("Avg is:%f", avg)

}
