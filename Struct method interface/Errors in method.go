// 10. Обработка ошибок в методах
// Задача: Создайте структуру Wallet с полем Balance. Добавьте метод Spend(amount float64) error,
// который уменьшает баланс, но возвращает ошибку, если денег недостаточно.
package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
}

func (w *Wallet) Spend(amount float64) error {
	w.Balance -= amount
	if w.Balance < 0 {
		return errors.New("Денег недостаточно")
	}
	return nil
}
func main() {
	Jack := Wallet{Balance: 100}
	Jack.Spend(15)
	Jack.Spend(30)
	fmt.Println(Jack.Balance)
	fmt.Println(Jack.Spend(70))
}
