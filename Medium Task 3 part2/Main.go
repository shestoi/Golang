// Создайте новую директорию с файлом main.go. Напишите код, в котором:
// нужно объединить слайс с выходными днями и слайс с рабочими в один слайс. Выведите на экран итоговый слайс с днями.
// Создайте новый коммит и отправьте его в удалённый репозиторий
package main

import "fmt"

func main() {
	workdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekend := []string{"Saturday", "Sunday"}
	for _, weekday := range weekend {
		workdays = append(workdays, weekday)
	}
	fmt.Println(workdays)
}
