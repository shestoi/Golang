// 6. Перебор мапы
// Дана мапа colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}.
// Выведите все ключи и значения в формате "red: #FF0000".
package main

import "fmt"

func main() {
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	for key, value := range colors {
		fmt.Println(key, ":", value)
	}
}
