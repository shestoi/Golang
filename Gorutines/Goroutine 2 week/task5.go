// 5. Банковский счет с RW Mutex
// Создайте структуру банковского счета, где:
// Проверка баланса происходит очень часто
// Изменение баланса (пополнение/списание) происходит редко
// Оптимизируйте с помощью sync.RWMutex.
package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	mu      sync.RWMutex
	balance int
}

func NewBankAccount(cash int) *BankAccount {
	return &BankAccount{
		balance: cash,
	}
}
func (account *BankAccount) Deposit(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.balance += amount
}
func (account *BankAccount) Balance() int {
	account.mu.RLock()
	defer account.mu.RUnlock()
	return account.balance
}
func main() {
	jack := NewBankAccount(100)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		jack.Deposit(10)
	}()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(jack.Balance())
		}()
	}
	wg.Wait()
}
