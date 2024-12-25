// 1)Создать пакет mathutils с функцией для вычисления факториала числа.
package main

import (
	"3Lab/mathutils"
	"3Lab/stringutils"
	"fmt"
)

func main() {
	//2)Использовать созданный пакет для вычисления факториала введенного пользователем числа.
	FirstNum := 5
	fmt.Printf("Factorial of  %d is %d \n", FirstNum, mathutils.Factorial(FirstNum))
	//3)Создать пакет stringutils с функцией для переворота строки и использовать его в основной программе.
	FirstString := "I love my mum"
	fmt.Printf("reverse string of %s\n is %s \n", FirstString, stringutils.Reversed(FirstString))
	//4)Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.
	var numbers [5]int
	for i := 0; i < 5; i++ {
		numbers[i] = i
		fmt.Printf("element:%d is %d\n", i+1, numbers[i])
	}
	//5)Создать срез из массива и выполнить операции добавления и удаления элементов.
	array := [5]int{11, 22, 33, 44, 55}
	slice := array[0:3]
	fmt.Println("Default slice:", slice)
	slice = append(slice, 66)
	fmt.Println("Slice after appending:", slice)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Slice after removing:", slice)
	//6)Написать программу, которая создает срез из строк и находит самую длинную строку.
	stringSlice := []string{"I", "love", "my", "mum"}
	fmt.Println("Longest string is ", stringutils.FindLongest(stringSlice))
}
