// 9. Паника в горутине
// Напишите программу, где одна из горутин вызывает panic, и обработайте её с помощью recover в главном потоке.

// В Go нельзя перехватить panic, возникшую в другой горутине, через recover() в главной. Это ограничение языка.

// package main
//
// import (
//
//	"fmt"
//	"sync"
//
// )
//
//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(2)
//		returnErr := make(chan interface{})
//		go func() {
//			defer wg.Done()
//			defer func() {
//				if err := recover(); err != nil {
//					returnErr <- err
//				}
//			}()
//			panic("panic in goroutine")
//		}()
//		go func() {
//			defer wg.Done()
//			fmt.Println("some goroutine")
//		}()
//		wg.Wait()
//
//		defer func() {
//
//			if returnErr <- recover(); returnErr != nil {
//				fmt.Println("В одной из горутин была паника")
//			}
//		}()
//	}
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	errCh := make(chan interface{}, 1) // канал для ошибок (panic)

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				errCh <- r // отправляем ошибку в главный поток
			}
		}()
		panic("что-то пошло не так в горутине")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("some goroutine")
	}()
	wg.Wait()
	close(errCh)

	if err, ok := <-errCh; ok {
		fmt.Println("В главной горутине получена паника из другой горутины:", err)
	} else {
		fmt.Println("Все горутины завершились без паники")
	}
}

//в качестве второго варианта мы обрабатываем панику внутри той же горутины
