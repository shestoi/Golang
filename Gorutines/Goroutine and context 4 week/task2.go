package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// 2. Параллельная обработка со сбором ошибок
// Реализуйте функцию, которая:
// Принимает слайс данных (чисел, строк и т.п.)
// Обрабатывает каждый элемент параллельно в горутинах
// Использует errgroup для управления горутинами
// Возвращает:
// Результаты обработки (если все элементы обработаны успешно)
// Первую возникшую ошибку (если хотя бы одна операция завершилась с ошибкой)
func processing(ctx context.Context, nums []int, strs []string) ([]int, []string, error) {
	g, ctx := errgroup.WithContext(ctx)
	resNums := []int{}
	resStrs := []string{}
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for _, v := range nums {
				if v < 0 {
					return errors.New("negative number")
				} else {
					resNums = append(resNums, v)
				}
			}
		}
		return nil
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for _, v := range strs {
				if v == "" {
					return errors.New("empty str")
				} else {
					resStrs = append(resStrs, v)
				}
			}
			return nil
		}
	})
	if err := g.Wait(); err != nil {
		return []int{}, []string{}, err
	}
	return resNums, resStrs, nil
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println(processing(ctx, []int{1, 2, 3, 4, 5}, []string{"a", "b", "c", "d"}))
}
