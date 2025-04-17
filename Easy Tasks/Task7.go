package main

import (
	"fmt"
	"golang.org/x/text/unicode/bidi"
	"strings"
)

func main() {
	var str string
	fmt.Println("Введите строку: ")
	fmt.Scanln(&str)
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	if str == bidi.ReverseString(str) {
		fmt.Println("Ваша строка палиндром")
	} else {
		fmt.Println("Ваша строка не палиндром")
	}
}
