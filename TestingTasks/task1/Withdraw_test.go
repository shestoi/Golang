package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

//	func TestBankAccount_WithdrawTableDriven(t *testing.T) {
//		var testcases = []struct {
//			text         string
//			input        float64
//			startBalance float64
//			wantErrType  interface{}
//			finalBalance float64
//		}{
//			{"Успешное снятие", 50, 100, nil, 50},
//			{"Недостаточно средств", 120, 100, &NotEnoughFundsError{}, 100},
//			{"Отрицательная сумма", -1, 100, &NegativeNumberError{}, 100},
//		}
//
//		for _, tt := range testcases {
//			t.Run(tt.text, func(t *testing.T) {
//				account := &BankAccount{balance: tt.startBalance}
//				err := account.Withdraw(tt.input)
//
//				switch want := tt.wantErrType.(type) {
//				case nil:
//					if err != tt {
//						t.Errorf("ожидалась nil, а пришла ошибка: %v", err)
//					}
//				default:
//					if !errors.As(err, &want) {
//						t.Errorf("ожидалась ошибка типа %T, а пришла: %v", want, err)
//					}
//				}
//
//				if account.GetBalance() != tt.finalBalance {
//					t.Errorf("баланс: ожидался %v, а получен %v", tt.finalBalance, account.GetBalance())
//				}
//			})
//		}
//	}
func TestBankAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name         string
		startBalance float64
		withdraw     float64
		wantErr      error // ожидаемый тип ошибки (не строка)
		wantBalance  float64
	}{
		{
			name:         "Успешное снятие",
			startBalance: 100,
			withdraw:     50,
			wantErr:      nil,
			wantBalance:  50,
		},
		{
			name:         "Недостаточно средств",
			startBalance: 100,
			withdraw:     120,
			wantErr:      &NotEnoughFundsError{},
			wantBalance:  100,
		},
		{
			name:         "Отрицательная сумма",
			startBalance: 100,
			withdraw:     -1,
			wantErr:      &NegativeNumberError{},
			wantBalance:  100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &BankAccount{balance: tt.startBalance}
			err := account.Withdraw(tt.withdraw)

			// Проверка ошибки
			if tt.wantErr != nil {
				require.Error(t, err)             //Прекращает тест, если ошибки нет, но она ожидалась
				assert.IsType(t, tt.wantErr, err) //Проверяет, что ошибка имеет нужный тип

			} else {
				require.NoError(t, err)
			}

			// Проверка баланса
			assert.Equal(t, tt.wantBalance, account.GetBalance()) //Проверяет значение баланса — даже если ошибка была

		})
	}
}
