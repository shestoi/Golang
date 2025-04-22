// 3. Удаление элемента по индексу
// Напишите функцию, которая удаляет элемент из среза по заданному индексу in-place (без возврата нового среза).
// Пример:
// Вход: slice = []int{10, 20, 30, 40}, index = 1
// Выход: [10 30 40]
package main

import "fmt"

func delete(ent *[]int, i int) {
	if i < 0 || i >= len(*ent) {
		fmt.Printf("ent[%d] is nil\n", i)
		return
	}
	*ent = append((*ent)[:i], (*ent)[i+1:]...)
}
func main() {
	ent := []int{10, 20, 30, 40}
	delete(&ent, 4)
	fmt.Println(ent)
}
