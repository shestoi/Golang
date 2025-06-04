// 4. Цепочка вызовов с ошибками: Реализуйте цепочку из 3 функций, где каждая может вернуть ошибку,
// и обрабатывайте ошибки на верхнем уровне.
package main

import (
	"errors"
	"fmt"
)

func sum(x int, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("отрицательный аргумент")
	}
	return x + y, nil
}
func getOperation(f func() (int, error)) {
	res, err := f()
	if err != nil {
		fmt.Println("Ошибка : ", err.Error())
	} else {
		fmt.Println(res)
	}
}
func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("деление на ноль")
	}
	return x / y, nil
}
func multiply(x int, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, errors.New("бессмысленная операция")
	}
	return x * y, nil
}
func main() {
	getOperation(func() (int, error) {
		return sum(13, -3)
	})
	getOperation(func() (int, error) {
		return divide(10, 0)
	})
	getOperation(func() (int, error) {
		return multiply(10, 0)
	})
}
