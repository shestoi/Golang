// 8. Graceful shutdown флаг
// Реализуйте механизм graceful shutdown, где:
// Одна горутина может установить флаг завершения
// Множество горутин могут проверять этот флаг
// Требуется максимальная производительность при проверках (atomic).
// package main
//
// import (
//
//	"fmt"
//	"sync"
//	"sync/atomic"
//
// )
//
//	func main() {
//		wg := sync.WaitGroup{}
//		flag := atomic.Bool{}
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				if flag.Load() {
//					fmt.Println("Flag is load")
//				} else {
//					fmt.Println("Flag is not load")
//				}
//			}()
//		}
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			flag.Store(true)
//		}()
//		wg.Wait()
//	}
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var flag atomic.Bool

	// Устанавливаем флаг в отдельной горутине чуть позже
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // Подождем, чтобы читающие запустились
		flag.Store(true)
		fmt.Println(">>> Флаг установлен в true")
	}()

	// 10 горутин читают флаг несколько раз
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				if flag.Load() {
					fmt.Printf("Горутина %d: флаг = true\n", id)
					return // Завершаем, если флаг активен
				}
				time.Sleep(50 * time.Millisecond)
			}
			fmt.Printf("Горутина %d: флаг так и не был установлен\n", id)
		}(i)
	}

	wg.Wait()
}
