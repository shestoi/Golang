// 4. Горутины и цикл
// Запустите 3 горутины внутри цикла, передавая каждой значение счётчика (i).
// Убедитесь, что вывод соответствует ожиданиям (0, 1, 2).
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go fmt.Println(i)
		time.Sleep(1 * time.Nanosecond)
	}
	time.Sleep(1 * time.Second)
}
