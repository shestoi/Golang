// 7. Фильтрация in-place
// Напишите функцию, которая удаляет из среза все четные числа in-place (без выделения нового среза).
// Пример:
// Вход: []int{1, 2, 3, 4, 5}
// Выход: [1 3 5]
package main

import "fmt"

func deleteEven(ent *[]int) {
	cpEnt := *ent
	inCp := 0
	for _, val := range *ent {
		if val%2 != 0 {
			cpEnt[inCp] = val
			inCp++
		}
	}
	*ent = cpEnt[:inCp]
}
func main() {
	ent := []int{1, 2, 3, 4, 5}
	deleteEven(&ent)
	fmt.Println(ent)
	ent2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //
	deleteEven(&ent2)
	fmt.Println(ent2)
}
