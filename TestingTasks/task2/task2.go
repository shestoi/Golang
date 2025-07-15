package main

import (
	"errors"
	"fmt"
)

//2. Условие:
//Ваша задача — модернизировать предыдущее решение и создать программу,
//которая демонстрирует использование интерфейсов для работы с различными типами банковских счетов.
//
//Требования:
//1.Создайте интерфейс Account :
//	Интерфейс должен содержать следующие методы:
//	GetBalance() float64: Возвращает текущий баланс счета.
//	Withdraw(amount float64) error: Снимает указанную сумму со счета (возвращает ошибку, если недостаточно средств).
//	CalculateTaxes() float64: Рассчитывает налоги на основе баланса счета.

type Account interface {
	GetBalance() float64           // Возвращает текущий баланс счета.
	Withdraw(amount float64) error // Снимает указанную сумму со счета (возвращает ошибку, если недостаточно средств).
	CalculateTaxes() float64       // Рассчитывает налоги на основе баланса счета.
}
type BankAccount2 struct {
	owner   string
	balance float64
}

//2. Реализуйте структуру TaxableBankAccount :
//Эта структура должна реализовывать интерфейс Account.
//Добавьте поле taxRate (процентная ставка налога), которое будет использоваться для расчета налогов.
//Метод CalculateTaxes() должен вычислять налог как баланс * taxRate / 100.

type TaxableBankAccount struct {
	BankAccount2
	taxRate float64 //процентная ставка налога
}

func (tb *TaxableBankAccount) GetBalance() float64 {
	return tb.balance
}
func (tb *TaxableBankAccount) Withdraw(amount float64) error {
	if tb.balance < amount {
		return errors.New("cannot withdraw bank account")
	} else {
		tb.balance -= amount
	}
	return nil
}
func (tb *TaxableBankAccount) CalculateTaxes() float64 { //должен вычислять налог как баланс * taxRate / 100.
	return tb.balance * tb.taxRate / 100
}

//3. Реализуйте структуру TaxableSavingsAccount :
//Эта структура также должна реализовывать интерфейс Account.
//Она должна быть основана на вашей уже существующей структуре SavingsAccount из предыдущей задачи.
//Для расчета налогов используйте фиксированную налоговую ставку (например, 3%).

type SavingsAccount2 struct {
	BankAccount2
	interestRate float64 //(процентная ставка)
}
type TaxableSavingsAccount struct {
	SavingsAccount2
}

func (ts *TaxableSavingsAccount) GetBalance() float64 {
	return ts.balance
}
func (ts *TaxableSavingsAccount) Withdraw(amount float64) error {
	if ts.balance < amount {
		return errors.New("cannot withdraw bank account")
	} else {
		ts.balance -= amount
	}
	return nil
}
func (ts *TaxableSavingsAccount) CalculateTaxes() float64 {
	return ts.balance * ts.interestRate / 100
}

//Пример использования программы:
//Предположим, у вас есть два счета:
//
//Банковский счет Ивана с балансом 3000 рублей и налоговой ставкой 13%.
//Сберегательный счет Ольги с балансом 5000 рублей и фиксированной налоговой ставкой 3%.
//Программа должна выводить результаты в следующем формате:
//Владелец: Иван, Баланс: 3000.00, Налоги: 390.00
//Владелец: Ольга, Баланс: 5000.00, Налоги: 150.00

func main() {
	ba := TaxableBankAccount{
		BankAccount2: BankAccount2{"Иван", 3000},
		taxRate:      13,
	}
	sa := TaxableSavingsAccount{
		SavingsAccount2{
			BankAccount2{"Ольга", 5000},
			150,
		},
	}
	fmt.Printf("%+v\n", ba)
	fmt.Printf("%+v\n", sa)
}

//Требования:
//Для каждого метода создайте таблицу тестовых случаев (table-driven tests)
//Каждый тестовый случай должен проверять конкретный сценарий
//Тесты должны проверять:
//Корректность работы методов при нормальных условиях
//Обработку граничных условий
//Возвращение ошибок в соответствующих случаях
//Используйте субтесты (t.Run) для лучшей читаемости и организации тестов
//
//Перепишите тесты для BankAccount и SavingsAccount, используя пакеты assert и require из библиотеки testify. Избегайте использования suite-тестирования.
//
//Требования:
//Для каждого метода создайте табличные тесты
//Используйте assert для проверок, которые могут продолжать тест после неудачи
//Используйте require для обязательных проверок (когда тест не имеет смысла продолжать)
//Для проверки ошибок используйте assert.Error, assert.NoError, assert.EqualError
//
//Конкретные задания:
//1. Перепишите тесты для Deposit:
//Проверьте положительные суммы
//Проверьте нулевую сумму
//Проверьте отрицательные суммы
//
//2. Перепишите тесты для Withdraw:
//Успешное снятие
//Недостаточно средств
//Отрицательная сумма
//Проверяйте точные тексты ошибок
//
//3. Перепишите тесты для ApplyInterest:
//Разные процентные ставки
//Нулевой баланс
//Проверка точности вычислений
//
//4. Добавьте тесты для проверки:
//Что SavingsAccount правильно встраивает BankAccount
//Что методы GetOwner и GetBalance работают через SavingsAccount
