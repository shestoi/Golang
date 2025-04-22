// 6. Объединение двух срезов
// Напишите функцию, которая принимает два среза и возвращает новый срез,
// содержащий все элементы обоих без использования append (на базе make и копирования).
// Пример:
// Вход: []int{1, 2}, []int{3, 4}
// Выход: [1 2 3 4]
package main

import "fmt"

func unification(first, second []int) []int {
	union := make([]int, len(first)+len(second))
	copy(union, first)
	copy(union[len(first):], second)
	return union
}
func main() {
	fmt.Println(unification([]int{1, 2}, []int{3, 4}))

}
