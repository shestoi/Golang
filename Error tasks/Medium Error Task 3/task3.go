package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Собственная ошибка
type LimitExceededError struct {
	message    string
	limit      int
	lastString string
}

func (e *LimitExceededError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
}

// Логика чтения файла и подсчета строк
func countLines(path string, limit int) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "ошибка при закрытии файла: %v\n", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	count := 0
	var last string

	for {
		if !scanner.Scan() {
			// Проверка на конец файла или другую ошибку
			if err := scanner.Err(); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return count, err
			}
			break
		}

		count++
		last = strings.TrimSpace(scanner.Text())

		if count > limit {
			return count, &LimitExceededError{
				message:    "Превышен лимит строк",
				limit:      limit,
				lastString: last,
			}
		}
	}

	return count, nil
}

// Основная функция
func main() {
	const limit = 3 // Произвольный лимит
	total, err := countLines("data/in2.txt", limit)

	var limErr *LimitExceededError
	if errors.As(err, &limErr) {
		fmt.Println("string count exceed limit, please read another file =)\nerr:", err.Error())
		return
	}

	if err != nil {
		fmt.Println("Неожиданная ошибка:", err)
		return
	}

	fmt.Printf("Total strings: %d\n", total)
}
