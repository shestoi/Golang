// 1. Простая горутина
// Напишите программу, которая запускает горутину, выводящую "Hello, Go!",
// и дожидается её завершения с помощью time.Sleep.
package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello, Go!")
	time.Sleep(1 * time.Second)
}
