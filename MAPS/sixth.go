// 6. Преобразование мапы в срез ключей/значений
// Напишите функцию, которая принимает map[int]string и возвращает два среза: один с ключами, другой со значениями.
// Пример:
// Вход: map[int]string{1: "a", 2: "b"}
// Выход: []int{1, 2}, []string{"a", "b"}
package main

import "fmt"

func divide(ent map[int]string) ([]int, []string) {
	keys := []int{}
	vals := []string{}
	for key, val := range ent {
		keys = append(keys, key)
		vals = append(vals, val)
	}
	return keys, vals
}
func main() {
	m := map[int]string{1: "a", 2: "b"}
	fmt.Println(divide(m))
}
