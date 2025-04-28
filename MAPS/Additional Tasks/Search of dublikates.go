// 9. Поиск дубликатов в слайсе
// Дан слайс nums := []int{1, 2, 3, 2, 4, 5, 4}.
// Проверьте, есть ли в нём дубликаты, используя мапу.
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 4}
	m := make(map[int]bool)
	dublicates := false
	for _, v := range nums {
		if !m[v] {
			m[v] = true
		} else {
			dublicates = true
		}
	}
	fmt.Println(dublicates)
}
