// Создайте новую директорию. В ней создайте файл main.go. Напишите код, в котором:
// из слайса с днями недели надо скопировать в новый слайс рабочие дни, а из исходного слайса удалить скопированные, чтобы остались только выходные дни.
// выведите на экран слайсы с выходными и рабочими днями недели.
package main

import "fmt"

func main() {
	var weekday = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var workday = weekday[:5]
	weekday = weekday[5:]
	fmt.Println("Work days : ", workday)
	fmt.Println("Weekends : ", weekday)
}
