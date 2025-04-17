package main

import "fmt"

func main() {
	var N int
	fmt.Println("Введите число: ")
	fmt.Scan(&N)
	fbnch := make([]int, N, N)
	if N < 0 {
		fmt.Println("Число Фибоначчи не может быть отрицательным")
	} else if N == 0 {
		fmt.Println(0)
	} else if N == 1 {
		fmt.Println(1)
	} else if N >= 2 {
		fbnch[0] = 1
		fbnch[1] = 2
		for i := 2; i < N; i++ {
			fbnch[i] = fbnch[i-1] + fbnch[i-2]
		}
	}
	fmt.Printf("%d число Фибоначчи равно : %d", N, fbnch[N-1])
}
