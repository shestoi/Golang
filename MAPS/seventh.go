// 7. Удаление дубликатов из среза с помощью мапы
// Напишите функцию, которая принимает срез строк и возвращает новый срез без дубликатов, используя мапу для проверки уникальности.
// Пример:
// Вход: []string{"a", "b", "a", "c"}
// Выход: []string{"a", "b", "c"}
package main

import "fmt"

func deleteDublicates(ent []string) []string {
	seen := make(map[string]bool)
	ex := []string{}
	for _, v := range ent {
		if !seen[v] {
			seen[v] = true
			ex = append(ex, v)
		}
	}
	return ex
}
func main() {
	fmt.Println(deleteDublicates([]string{"a", "b", "a", "c"}))
}
