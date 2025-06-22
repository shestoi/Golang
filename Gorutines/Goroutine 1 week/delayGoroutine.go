// 3. Горутина с задержкой
// Создайте горутину, которая ждёт 2 секунды (time.Sleep), а затем выводит "Done".
// Главный поток должен дождаться её завершения.
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done")
	}()
	time.Sleep(3 * time.Second)

}
