// 3. Проверка типа ошибки: Напишите код, который пытается преобразовать строку в число и по-разному
// обрабатывает ошибки парсинга и ошибки диапазона.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type ParsingError struct{}

func (parsingError ParsingError) Error() string {
	return "parsing Error"
}

type DiapasonError struct{}

func (diapasonError DiapasonError) Error() string {
	return "diapasonError"
}
func Transformation(str string) (int, error) { //в этой функции мы пытаемся сделать наше преобразование
	result, error := strconv.Atoi(str) // и смотрим какие ошибки могут возникунуть
	if error != nil {
		return 0, ParsingError{}
	}
	if result < 0 { //за ошибку диапазона мы считаем отрицательные значения
		return 0, DiapasonError{}
	}
	return result, nil
}
func BT(str string) { //Здесь мы красиво оформляем нашу итоговую картинку
	_, err := Transformation(str) //функции траснформация
	if errors.Is(err, ParsingError{}) {
		fmt.Println(ParsingError{})
	} else if errors.Is(err, DiapasonError{}) {
		fmt.Println(DiapasonError{})
	} else {
		res, _ := Transformation(str)
		fmt.Println(res)
	}

}
func main() {
	BT("abc")
	BT("-111")
	BT("420")
}
