package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Открываем файл
	file, err := os.Open("data/in2.txt")
	if err != nil {
		panic(fmt.Sprintf("Ошибка при открытии файла: %v", err))
	}
	// Закрываем файл с обработкой ошибки
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при закрытии файла: %v\n", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Читаем построчно
	for scanner.Scan() {
		lineCount++
	}

	// Проверяем наличие ошибок при сканировании (например, EOF — это не ошибка, но другие могут быть)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении файла: %v\n", err)
	}

	fmt.Printf("Total strings: %d\n", lineCount)
}
