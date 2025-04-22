// 5. Передача среза по указателю
// Напишите функцию, которая принимает указатель на срез строк и добавляет строку "modified" в конец.
// Пример:
// Вход: &[]string{"a", "b"}
// Выход: ["a" "b" "modified"]
package main

import "fmt"

func modified(ent *[]string) {
	*ent = append(*ent, "modified")
}
func main() {
	ent := []string{"a", "b"}
	modified(&ent)
	fmt.Println(ent)
}
