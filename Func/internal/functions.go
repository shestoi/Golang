package internal

import "errors"

//1. Подсчёт частоты элементов в слайсе
//Напишите функцию CountFrequency, которая принимает слайс строк и возвращает мапу,
//где ключи — элементы слайса, а значения — количество их вхождений.

func CountFrequency(slice []string) map[string]int {
	count := map[string]int{}
	for _, v := range slice {
		count[v]++
	}
	return count
}

//2. Фильтрация слайса по условию
//Напишите функцию FilterEven, которая принимает слайс целых чисел и возвращает новый слайс,
//содержащий только чётные числа.

func FilterEven(slice []int) []int {
	even := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

// 3. Объединение двух мап
// Напишите функцию MergeMaps, которая принимает две мапы (map[string]int) и возвращает новую мапу,
// объединяющую их. Если ключ присутствует в обеих мапах, выбирается значение из второй.

func MergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// 4. Удаление дубликатов из слайса
// Напишите функцию RemoveDuplicates,
// которая принимает слайс строк и возвращает новый слайс без дубликатов.

func RemoveDuplicates(s []string) []string {
	result := []string{}
	seen := map[string]bool{}
	for _, str := range s {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

//5. Напишите функцию InvertMap, которая принимает мапу map[string]int
//и возвращает новую мапу map[int]string, где ключи и значения поменялись местами.
//Если в исходной мапе есть дублирующиеся значения, вернуть ошибку.

func InvertMap(ent map[string]int) (map[int]string, error) {
	result := map[int]string{}
	seen := map[int]bool{}
	for k, v := range ent {
		if !seen[v] {
			seen[v] = true
			result[v] = k
		} else {
			return result, errors.New("duplicate map")
		}
	}
	return result, nil
}

//6. Разделение слайса на пакеты (чанки)
//Напишите функцию ChunkSlice, которая принимает слайс целых чисел и размер чанка,
//а возвращает слайс слайсов, разбитый на части указанного размера.

func ChunkSlice(slice []int, chunkSize int) [][]int {
	result := [][]int{}
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

//7. Поиск уникальных элементов в слайсе
//Напишите функцию FindUnique, которая принимает слайс целых чисел и возвращает новый слайс,
//содержащий только уникальные элементы (встречающиеся 1 раз).

func FindUnique(slice []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, v := range slice {
		seen[v]++
	}
	for k, v := range seen {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

//8. Подсчёт среднего значения в слайсе
//Напишите функцию Average, которая принимает слайс чисел
//и возвращает их среднее значение. Если слайс пуст, возвращает 0.

// func Average(slice []int) float64 {
func Average(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	//average := float64(0)
	//for _, v := range slice {
	//	average += float64(v)
	//}
	//return average/float64(len(slice))
	average := 0
	for _, v := range slice {
		average += v
	}
	return (average) / len(slice)
}

//9. Группировка строк по длине
//Напишите функцию GroupByLength, которая принимает слайс строк и возвращает мапу,
//где ключ — длина строки, а значение — слайс строк этой длины.

func GroupByLength(slice []string) map[int][]string {
	result := map[int][]string{}
	for _, v := range slice {
		result[len(v)] = append(result[len(v)], v)
	}
	return result
}
