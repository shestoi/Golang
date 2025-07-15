package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 4. Добавьте тесты для проверки:
// Что SavingsAccount правильно встраивает BankAccount
// Что методы GetOwner и GetBalance работают через SavingsAccount
func TestBankAccount_StructsEmbeddingTableDriven(t *testing.T) {
	tests := []struct {
		owner        string
		balance      float64
		interestRate float64
		expectedBal  float64
	}{
		{"Mary", 120, 20, 144},
		{"Henry", 0, 10, 0},
		{"Erik", 2000, 15, 2300},
	}
	for _, tt := range tests {
		t.Run(tt.owner, func(t *testing.T) {
			sa := SavingsAccount{BankAccount{tt.owner, tt.balance}, tt.interestRate}
			assert.Equal(t, tt.owner, sa.GetOwner(), "GetOwner должен вернуть имя владельца")
			assert.Equal(t, tt.balance, sa.GetBalance(), "GetBalance должен вернуть текущий баланс")

			assert.Equal(t, tt.interestRate, sa.interestRate, "interestRate должен совпадать")
		})
	}
}
