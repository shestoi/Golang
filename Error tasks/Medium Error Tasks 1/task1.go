package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла паника:", r)
			// читаем, что записали до ошибки
			fmt.Println("\nСодержимое data_out.txt до ошибки:")
			content, err := os.ReadFile("data/data_out.txt") //
			if err != nil {
				fmt.Println("Не удалось открыть файл data_out.txt:", err)
			} else {
				fmt.Println(string(content))
			}
		}
	}()

	inFile, err := os.Open("data/in1.txt")
	if err != nil {
		panic("не удалось открыть data/in1.txt: " + err.Error())
	}
	defer inFile.Close()

	outFile, err := os.Create("data/data_out.txt")
	if err != nil {
		panic("не удалось создать data_out.txt: " + err.Error())
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile) //создали сканнер который будет считывать строки из inFile
	row := 0

	for scanner.Scan() { //создали цикл который будет работать пока в сканнере есть строки
		row++
		line := scanner.Text()            //получаем текст текущей строки
		parts := strings.Split(line, "|") //разделяем строки на части через '|'

		if len(parts) != 3 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" || strings.TrimSpace(parts[2]) == "" {
			panic(fmt.Sprintf("parse error: empty field on string %d", row))
		}
		//проверили что в каждой строке 3 символа, и каждый из них не пуст
		formatted := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n", row, parts[0], parts[1], parts[2]) //создали строку которая требуется по шаблону
		outFile.WriteString(formatted)                                                                              //записали ее
	}
}
