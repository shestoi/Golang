// 3. Проверка наличия ключа
// Дана мапа ages := map[string]int{"Alice": 25, "Bob": 30}.
// Проверьте, есть ли в ней ключ "Bob", и выведите "Exists" или "Not exists"
package main

import "fmt"

func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	//if _, ok := ages["Keny"]; ok {
	//	fmt.Println("Exists")
	//} else {
	//	fmt.Println("Not Exists")
	//}
	_, ok := ages["Bob"]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not Exists")
	}
}
