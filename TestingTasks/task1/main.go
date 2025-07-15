package main

import (
	"fmt"
)

type NegativeNumberError struct{}

func (e *NegativeNumberError) Error() string {
	return "negative number error"
}

type NotEnoughFundsError struct{}

func (e *NotEnoughFundsError) Error() string {
	return "not enough funds"
}

//1. Условие:
//Реализуйте структуру BankAccount с полями:
//
//owner (владелец, приватное поле)
//balance (баланс, приватное поле)
//
//Добавьте методы:
//	Deposit(amount float64) – пополнение счета
//	Проверяет что amount > 0 и пополняет счет
//	Withdraw(amount float64) error – снятие (возвращает ошибку, если недостаточно средств)
//	GetBalance() float64 – текущий баланс
//	GetOwner() string – имя владельца
//
//Создайте структуру SavingsAccount,представляет сберегательный счет с начислением процентов.
//Это структура которая встраивает BankAccount и добавляет:
//
//interestRate (процентная ставка)
//метод ApplyInterest() – начисляет проценты на баланс

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}
func (b *BankAccount) Withdraw(amount float64) error { //снятие (возвращает ошибку, если недостаточно средств)
	if b.balance < amount {
		return &NotEnoughFundsError{}
	} else if amount < 0 {
		return &NegativeNumberError{}
	} else {
		b.balance -= amount
	}
	return nil
}
func (b *BankAccount) GetBalance() float64 { //текущий баланс
	return b.balance
}
func (b *BankAccount) GetOwner() string {
	return b.owner
}

type SavingsAccount struct {
	BankAccount
	interestRate float64 //(процентная ставка)
}

func (s *SavingsAccount) ApplyInterest() { //– начисляет проценты на баланс
	s.balance += s.balance * s.interestRate / 100
}
func main() {
	b := BankAccount{
		owner:   "Marat",
		balance: 100,
	}
	var bankAccount = SavingsAccount{
		b,
		2,
	}
	//b2 := SavingsAccount{
	//	BankAccount{"asd", 100},
	//	123,
	//}
	bankAccount.ApplyInterest()
	fmt.Println(bankAccount.GetBalance())

}
