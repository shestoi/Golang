// 5. Вложенные мапы (словарь словарей)
// Создайте map[string]map[string]int, где внешний ключ — имя студента, а внутренняя мапа хранит предметы и оценки. Напишите функцию для добавления оценки студенту.
// Пример:
// Вход:
//
// db := map[string]map[string]int{}
// addGrade(db, "Alice", "Math", 90)
// Выход: map[Alice:map[Math:90]]
package main

import "fmt"

func addGrade(name, lesson string, grade int) map[string]map[string]int {
	j := map[string]map[string]int{}
	j[name] = map[string]int{}
	j[name][lesson] = grade
	return j
}
func main() {
	fmt.Println(addGrade("Alice", "Math", 90))
}
