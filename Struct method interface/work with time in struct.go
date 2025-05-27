// 16. Работа с временем в структуре
// Задача:
// Создайте структуру Event с полями Title и Time (типа time.Time). Добавьте метод
// IsAfter(now time.Time) bool, проверяющий, что событие еще не наступило.
package main

import (
	"fmt"
	"time"
)

type Event struct {
	Title string
	Time  time.Time
}

func (e Event) IsAfter(now time.Time) bool {
	if e.Time.After(now) {
		return true
	} else {
		return false
	}
}
func main() {
	e := Event{Title: "Hello World", Time: time.Date(2026, time.May, 16, 14, 45, 0, 0, time.Local)}
	fmt.Println(e.IsAfter(time.Now()))
	Birthday := Event{Title: "Birthday", Time: time.Date(2002, time.June, 9, 8, 5, 0, 0, time.Local)}
	fmt.Println(Birthday.IsAfter(time.Now()))
}
