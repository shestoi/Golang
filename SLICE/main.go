// 1.Напишите функцию, которая создает срез целых чисел от 1 до n,
// затем добавляет к нему числа от n+1 до 2n с помощью append.
// Выведите исходный и измененный срез.
//
// Пример:
// Вход: n = 3
// Выход:
// Исходный: [1 2 3]
// После append: [1 2 3 4 5 6]
package main

import "fmt"

func getSlice(n int) {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	fmt.Println("Исходный: ", s)
	for i := n + 1; i <= 2*n; i++ {
		s = append(s, i)
	}
	fmt.Println("После append: ", s)
}
func main() {
	getSlice(3)
}
