package main

import (
	"testing"
)

// //1. Перепишите тесты для Deposit:
// //Проверьте положительные суммы
// //Проверьте нулевую сумму
// //Проверьте отрицательные суммы
// //
func TestBankAccount_DepositTableDriven(t *testing.T) {
	var testcases = []struct {
		text        string
		input       float64
		wantBalance float64
	}{
		{"-1 - отрицательное число", -1, 0},
		{"0 - нулевое число", 0, 0},
		{"1 - положительное число", 1, 1},
	}

	for _, tt := range testcases {
		t.Run(tt.text, func(t *testing.T) {
			account := &BankAccount{} // создаём аккаунт
			account.Deposit(tt.input)

			if account.balance != tt.wantBalance {
				t.Errorf("got balance %f, want %f", account.balance, tt.wantBalance)
			}
		})
	}
}
