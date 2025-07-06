package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// 4. Построение конвейеров обработки данных
// Реализуйте конвейер (pipeline), где:
// Генератор создаёт поток данных (например, числа от 1 до 100)
// Фильтр пропускает только данные, удовлетворяющие условию (например, чётные числа)
// Обработчик преобразует данные (например, умножает на 10)
// Все этапы должны поддерживать отмену через context.Context
// При ошибке на любом этапе весь конвейер останавливается

func generate(ctx context.Context, out chan<- int) error {
	defer close(out)
	for i := 1; i <= 100; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- i:
		}
	}
	return nil
}

func filter(ctx context.Context, in <-chan int, out chan<- int) error {
	defer close(out)
	for val := range in {
		if val == 80 {
			return fmt.Errorf("ошибка в фильтрации")
		}
		if val%2 != 0 {
			continue
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- val:
		}
	}
	return nil
}

func processor(ctx context.Context, in <-chan int, out chan<- int) error {
	defer close(out)
	for val := range in {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- val * 10:
		}
	}
	return nil
}

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	genChan := make(chan int)
	filtChan := make(chan int)
	procChan := make(chan int)

	g.Go(func() error {
		return generate(ctx, genChan)
	})
	g.Go(func() error {
		return filter(ctx, genChan, filtChan)
	})
	g.Go(func() error {
		return processor(ctx, filtChan, procChan)
	})

	// Вывод в отдельной горутине, не блокирует errgroup
	g.Go(func() error {
		for val := range procChan {
			fmt.Println("Результат:", val)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Конвейер завершён с ошибкой:", err)
	} else {
		fmt.Println("Конвейер успешно завершён")
	}
}
