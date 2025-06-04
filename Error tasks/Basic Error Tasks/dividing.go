// 1. Напишите функцию, которая делит два числа и возвращает ошибку, если делитель равен нулю.
package main

import (
	"errors"
	"fmt"
)

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func main() {
	a := 12
	b := 3
	c := 0
	fmt.Println(Div(a, b))
	fmt.Println(Div(a, c))
}
