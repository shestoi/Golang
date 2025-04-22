// 4. Проверка на анаграммы
// Напишите функцию, которая проверяет, являются ли две строки анаграммами (содержат одинаковые символы в разном порядке), используя мапу для подсчета символов.
// Пример:
// Вход: "listen", "silent"
// Выход: true
package main

import "fmt"

func isAnagram(s string, t string) bool {
	ms := make(map[rune]int)
	mt := make(map[rune]int)
	for _, r := range s {
		ms[r]++
	}
	for _, r := range t {
		mt[r]++
	}
	if len(ms) != len(mt) {
		return false
	}
	for k, v := range mt {
		if ms[k] != v {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram("listen", "silent"))
}
