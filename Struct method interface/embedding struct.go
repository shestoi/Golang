// 4. Встраивание структур
// Задача: Создайте структуру Person с полями Name и Age.
// Затем создайте структуру Employee, встраивающую Person, и добавляющую поле Salary.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//создали структуру, добавили поля

type Employee struct {
	Person //указали на то что у него присутствуют поля Person
	Salary int
}

func main() {
	e := Employee{
		Person: Person{ //Нам нужно заполнить поля Person
			Name: "John Doe",
			Age:  40,
		}, //После заполнения Person - ставим запятую, так как заполнение
		Salary: 1000, //Полей Employee еще не закончилось
	}
	fmt.Println(e)
}
