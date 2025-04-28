// 10. Инвертирование мапы
// Дана мапа m := map[int]string{1: "one", 2: "two"}.
// Создайте новую мапу, где ключи и значения поменяются местами (map[string]int).
package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two"}
	mi := make(map[string]int)
	for k, v := range m {
		mi[v] = k
	}
	fmt.Println(mi)
}
