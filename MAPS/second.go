// 2. Объединение двух мап
// Напишите функцию, которая принимает две map[string]int и возвращает новую мапу, объединяя их. Если ключ присутствует в обеих мапах, сложите значения.
// Пример:
// Вход:
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Выход: map[a:1 b:5 c:4]
package main

import "fmt"

func unification(first, second map[string]int) map[string]int {
	union := map[string]int{}
	for k, v := range first {
		union[k] = v
	}

	for k, v := range second {
		if union[k] > 0 {
			union[k] = union[k] + v
		} else {
			union[k] = v
		}
	}
	return union
}
func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	fmt.Println(unification(m1, m2))
}
