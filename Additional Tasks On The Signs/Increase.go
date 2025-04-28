// 2. Увеличение числа через указатель
// Задача:
// Напишите функцию Increment(n *int), которая увеличивает значение переданного числа на 1.
//
// Пример:
//
// num := 5
// Increment(&num)
// fmt.Println(num) // 6
// Условие:
//
// Функция не должна возвращать значение, а изменять переданную переменную.
package main

import "fmt"

func Increment(n *int) {
	*n++
}
func main() {
	num := 5
	Increment(&num)
	fmt.Println(num) // 6
}
