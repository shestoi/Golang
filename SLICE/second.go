// 2. Изменение in-place
// Напишите функцию, которая принимает срез целых чисел и умножает каждый его элемент на 2 без возврата нового среза (изменя исходный).
// Пример:
// Вход: []int{1, 2, 3}
// Выход: [2 4 6]
package main

import "fmt"

func multiplication(ent []int) {
	for i := range ent {
		(ent)[i] *= 2
	}
}
func main() {
	ent := []int{1, 2, 3}
	multiplication(ent)
	fmt.Println(ent)
}
