// 3. Фильтрация мапы по значению
// Напишите функцию, которая принимает map[string]int и возвращает новую мапу, содержащую только элементы со значениями больше заданного числа n.
// Пример:
// Вход: map[string]int{"a": 10, "b": 5, "c": 20}, n = 15
// Выход: map[c:20]
package main

import "fmt"

func filter(ent map[string]int, n int) map[string]int {
	m := make(map[string]int)
	for k, v := range ent {
		if v > n {
			m[k] = v
		}
	}
	return m
}
func main() {
	fmt.Println(filter(map[string]int{"a": 10, "b": 5, "c": 20}, 15))

}
