package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// 1. Таймауты и отмена операций
// Напишите функцию, которая:
// Выполняет долгую операцию (например, time.Sleep или цикл вычислений)
// Принимает context.Context для контроля выполнения
// Прекращает работу, если контекст отменён (например, по таймауту 2 секунды)
// Возвращает ошибку, если операция не завершилась вовремя
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	fmt.Println(worker(ctx))
}
func worker(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		time.Sleep(time.Second * 3)
		select {
		case <-ctx.Done():
			return nil
			//return ctx.Err()
		default:
			fmt.Println("some work")
			return nil
		}
	})
	g.Go(func() error {
		fmt.Println("second started")
		return fmt.Errorf("операция не успела завершится")
	})
	g.Go(func() error {
		fmt.Println("third started")
		return nil
	})
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
