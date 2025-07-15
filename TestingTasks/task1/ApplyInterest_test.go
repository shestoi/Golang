package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSavingsAccount_ApplyInterest(t *testing.T) {
	tests := []struct {
		name         string
		initialBal   float64
		interestRate float64
		expectedBal  float64
	}{
		{
			name:         "positive interest rate",
			initialBal:   1000,
			interestRate: 5,
			expectedBal:  1050,
		},
		{
			name:         "zero interest rate",
			initialBal:   1000,
			interestRate: 0,
			expectedBal:  1000,
		},
		{
			name:         "zero balance",
			initialBal:   0,
			interestRate: 10,
			expectedBal:  0,
		},
		{
			name:         "fractional interest rate",
			initialBal:   1234.56,
			interestRate: 3.5,
			expectedBal:  1277.7696,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем аккаунт с начальными данными
			account := SavingsAccount{
				BankAccount:  BankAccount{balance: tt.initialBal},
				interestRate: tt.interestRate,
			}

			// Применяем проценты
			account.ApplyInterest()

			// require - если баланс не совпадет, тест дальше смысла нет
			require.InDelta(t, tt.expectedBal, account.balance, 0.0001, "Баланс после начисления процентов отличается")

			// assert - можно добавить дополнительные проверки, если нужно
			assert.True(t, account.balance >= 0, "Баланс должен быть неотрицательным")
		})
	}
}
