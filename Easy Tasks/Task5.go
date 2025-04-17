package main

import (
	"fmt"
	"golang.org/x/text/unicode/bidi"
)

func main() {
	var str string
	fmt.Print("Введите строку которую хотите перевернуть: ")
	fmt.Scanln(&str)
	fmt.Printf("Ваша строка в обратном порядке: %s", bidi.ReverseString(str))
}
