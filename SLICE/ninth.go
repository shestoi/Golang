// 9. Вставка элемента по индексу
// Напишите функцию, которая вставляет элемент в срез по указанному индексу с сохранением исходного capacity.
// Пример:
// Вход: slice = []int{10, 20, 30}, index = 1, value = 99
// Выход: [10 99 20 30]
package main

import "fmt"

func insertValue(ent *[]int, i, value int) {
	*ent = append((*ent)[:i], append([]int{value}, (*ent)[i:]...)...)
}
func main() {
	ent := make([]int, 5, 5)

	ent = []int{10, 20, 30}
	fmt.Println(len(ent))
	insertValue(&ent, 1, 99)
	fmt.Println(len(ent))
	fmt.Println(ent)
}
