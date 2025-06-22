// 2. Множество горутин
// Запустите 5 горутин, каждая из которых выводит свой номер (от 1 до 5).
// Убедитесь, что все горутины успевают выполниться.
package main

import (
	"fmt"
	"time"
)

// переделать под циклы
func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println(i, "goroutine")
		time.Sleep(1 * time.Second)
	}
	//go fmt.Println("First Goroutine")
	//time.Sleep(1 * time.Second)
	//go fmt.Println("Second Goroutine")
	//time.Sleep(1 * time.Second)
	//go fmt.Println("Third Goroutine")
	//time.Sleep(1 * time.Second)
	//go fmt.Println("Fourth Goroutine")
	//time.Sleep(1 * time.Second)
	//go fmt.Println("Fifth Goroutine")
	//time.Sleep(1 * time.Second)
}
