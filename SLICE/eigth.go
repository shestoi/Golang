// 8. Разворот среза
// Напишите функцию, которая разворачивает срез in-place (первый элемент становится последним и т.д.).
// Пример:
// Вход: []int{1, 2, 3, 4}
// Выход: [4 3 2 1]
package main

import "fmt"

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4}
	reverseSlice(s)
	fmt.Println(s)
}
