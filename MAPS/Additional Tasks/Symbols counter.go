// 8. Подсчёт частоты символов в строке
// Дана строка "hello".
// Создайте мапу, где ключи — символы, а значения — сколько раз они встречаются.
package main

import "fmt"

func main() {
	str := "hello"
	m := make(map[rune]int)
	for _, v := range str {
		m[v]++
	}
	for k, v := range m {
		fmt.Println(string(k), ":", v)
	}
}
