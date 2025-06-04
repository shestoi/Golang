// 2. Создание кастомных ошибок: Создайте тип NegativeNumberError и напишите функцию, которая вычисляет
// квадратный корень и возвращает эту ошибку, если число отрицательное.
package main

import (
	"fmt"
	"math"
)

type NegativeNumberError struct{}

func (e NegativeNumberError) Error() string {

	return "Ошибка, число отрицательное"
}
func Sqrt(x int) (int, error) {
	if x < 0 {
		return 0, NegativeNumberError{}
	}
	return int(math.Sqrt(float64(x))), nil
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
