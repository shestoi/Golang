// 5. Гонка данных
// Создайте программу с общей переменной, которую увеличивают несколько горутин.
// Покажите, что без синхронизации возможна гонка данных.
package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	go func() {
		i++
		fmt.Println("first goroutine", i)
	}()
	go func() {
		i++
		fmt.Println("second goroutine", i)
	}()
	go func() {
		i++
		fmt.Println("third goroutine", i)
	}()
	time.Sleep(1 * time.Second)
}
