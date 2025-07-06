package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 3. Управление жизненным циклом горутин
// Создайте программу, которая:
// Запускает N рабочих горутин (например, 3)
// Каждая горутина периодически выполняет свою задачу
// При получении сигнала SIGINT (Ctrl+C):
// Корректно останавливает все горутины
// Даёт им время на завершение текущих операций
// Выводит статус завершения
func main() {
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	defer stop()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 1 получила сигнал отмены - завершаем")
				return
			default:
				fmt.Println("Горутина 1 работает")
				time.Sleep(2 * time.Second)

			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 2 получила сигнал отмены - завершаем")
				return
			default:
				fmt.Println("Горутина 2 работает")
				time.Sleep(3 * time.Second)

			}
		}
	}()
	wg.Wait()
	fmt.Println("Завершаем")
}
