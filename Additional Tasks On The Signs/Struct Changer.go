// 5. Изменение структуры через указатель
// Задача:
// Дана структура:
//
//	type Person struct {
//	   Name string
//	   Age  int
//	}
//
// Напишите функцию Birthday(p *Person), которая увеличивает возраст (Age) на 1.
//
// Пример:
//
// person := Person{Name: "Alice", Age: 25}
// Birthday(&person)
// fmt.Println(person.Age) // 26
// Условие:
//
// Изменять нужно исходную структуру, а не возвращать новую.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Birthday(p *Person) {
	p.Age++
}
func main() {
	person := Person{Name: "Alice", Age: 25}
	Birthday(&person)
	fmt.Println(person.Age) // 26
}
