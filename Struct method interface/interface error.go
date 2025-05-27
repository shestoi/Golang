// 13. Реализация интерфейса error
// Задача:
// Создайте структуру DivError с полями dividend и divisor. Реализуйте для нее интерфейс error,
// чтобы при делении на 0 возвращалось сообщение: "деление на 0: dividend={X}, divisor=0".
package main

import "fmt"

type DivError struct {
	dividend, divisor int
}
type error interface {
	Error() string
}

func (e DivError) Error() string {
	if e.divisor == 0 {
		return fmt.Sprintf("деление на 0: dividend={%d}, divisor=0", e.dividend)
	} else {
		return "" //
	}
}

func main() {
	n1 := DivError{
		divisor:  2,
		dividend: 3,
	}
	n2 := DivError{
		divisor:  0,
		dividend: 2,
	}
	fmt.Println(error(n1))
	fmt.Println(error(n2))
}
