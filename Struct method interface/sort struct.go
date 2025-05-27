// 7. Сортировка структур
// Задача: Создайте структуру Student (Name, Grade).
// Создайте слайс студентов и отсортируйте его по полю Grade (используя sort.Interface).
package main

import (
	"fmt"
	"sort"
)

// Структура Student
type Student struct {
	Name  string
	Grade int
}

// Слайс студентов, реализующий sort.Interface
type ByGrade []Student

// Методы интерфейса sort.Interface:

// Длина слайса
func (s ByGrade) Len() int {
	return len(s)
}

// Условие, по которому идёт сортировка (по возрастанию Grade)
func (s ByGrade) Less(i, j int) bool {
	return s[i].Grade < s[j].Grade
}

// Меняем местами элементы
func (s ByGrade) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
		{"Diana", 90},
	}

	fmt.Println("До сортировки:")
	for _, s := range students {
		fmt.Println(s.Name, s.Grade)
	}

	// Сортировка по Grade
	sort.Sort(ByGrade(students))

	fmt.Println("\nПосле сортировки:")
	for _, s := range students {
		fmt.Println(s.Name, s.Grade)
	}
}
