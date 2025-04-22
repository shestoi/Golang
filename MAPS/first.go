// 1. Подсчет частоты элементов
// Напишите функцию, которая принимает срез строк и возвращает map[string]int, где ключ — это строка, а значение — количество её вхождений.
// Пример:
// Вход: []string{"a", "b", "a", "c", "b"}
// Выход: map[a:2 b:2 c:1]
package main

import "fmt"

func count(ent []string) map[string]int {
	m := make(map[string]int)
	for _, e := range ent {
		m[e]++
	}
	return m
}
func main() {
	fmt.Println(count([]string{"a", "b", "a", "c", "b"}))
}
