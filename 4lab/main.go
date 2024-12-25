package main

import (
	"fmt"
	"strings"
)

func AverageAge(humanMap map[string]int) float64 {
	var totalAge int = 0
	var counter int = 0
	for _, value := range humanMap {
		counter++
		totalAge += value
	}
	return float64(totalAge) / float64(counter)
}

func Output(humanMap map[string]int) {
	for key, value := range humanMap {
		fmt.Println("Name:", key, "Age:", value)
	}
}
func EnterPeople(humanMap map[string]int) {
	buffName := ""
	buffAge := 0
	fmt.Println("Enter people in format 'name age'")
	for {

		fmt.Scan(&buffName, &buffAge)
		if buffName == "q" {
			break
		}
		humanMap[buffName] = buffAge
	}
	Output(humanMap)
}
func DeletePeople(humanMap map[string]int) {
	buffName := ""
	fmt.Println("Enter name to delete")
	fmt.Scan(&buffName)
	delete(humanMap, buffName)
	fmt.Println("our map")
	Output(humanMap)
}
func ToUp() {
	fmt.Println("Введите строку: ")
	var inputString string
	fmt.Scan(&inputString)
	inputString = strings.ToUpper(inputString)
	fmt.Println(inputString)
}
func Sum() {
	fmt.Println("enter nums to add or 0 to finish")
	var sum float64
	var buffNum float64
	for {
		fmt.Scan(&buffNum)
		sum += buffNum
		if buffNum == 0 {
			break
		}
	}
	fmt.Println(sum)
}
func main() {
	//1)Написать программу, которая создает карту с именами людей и их возрастами. Добавить нового человека и вывести все записи на экран.
	var humanMap map[string]int
	humanMap = map[string]int{
		"Pavel":      00,
		"Konstantin": 20,
		"Bill":       666,
	}
	EnterPeople(humanMap)
	// 2)Реализовать функцию, которая принимает карту и возвращает средний возраст всех людей в карте.
	fmt.Printf("average age is %f\n", AverageAge(humanMap))
	//3)Написать программу, которая удаляет запись из карты по заданному имени.
	DeletePeople(humanMap)
	//4)Написать программу, которая считывает строку с ввода и выводит её в верхнем регистре.
	ToUp()
	//5)Написать программу, которая считывает несколько чисел, введенных пользователем, и выводит их сумму.
	Sum()
	//6)Написать программу, которая считывает массив целых чисел и выводит их в обратном порядке.
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("%d ", array[i])
	}
}
