// 14. Метод с variadic-аргументами
// Задача:
// Создайте структуру Calculator с методом Sum(numbers ...int) int,
// который возвращает сумму переданных чисел.
package main

import "fmt"

type Calculator struct {
}

func (c Calculator) Sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
func main() {
	c := Calculator{}
	fmt.Println(c.Sum(1, 2, 3))
}
