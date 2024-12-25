/*

6)Создать интерфейс Stringer и реализовать его для структуры Book, которая хранит информацию о книге.*/

package main

import (
	"fmt"
	"math"
)

// 1)Создать структуру Person с полями name и age. Реализовать метод для вывода информации о человеке.
type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Printf("Name: %s, age: %d years old.\n", p.name, p.age)
}

// 2)Реализовать метод birthday для структуры Person, который увеличивает возраст на 1 год.
func (p Person) Birthday() {
	p.age++
	fmt.Printf("Happy birthday %s! Now you are %d years old.\n", p.name, p.age)
}

// 3)Создать структуру Circle с полем radius и метод для вычисления площади круга.
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	var area = c.radius * c.radius * math.Pi
	return area
}

// 4)Создать интерфейс Shape с методом Area(). Реализовать этот интерфейс для структур Rectangle и Circle.
type Shape interface {
	Area() float64
}
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	var area = r.length * r.width
	return area
}
func Area(shape Shape) {
	fmt.Println(shape.Area())
}

// 5)Реализовать функцию, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта.
func EveryArea(shapes []Shape) {
	for range shapes {
		fmt.Println(shapes[0].Area())
	}
}
func main() {
	person := Person{name: "boba", age: 30}
	person.Info()
	person.Birthday()
	circle := Circle{radius: 5}
	Shape.Area(circle)

}
