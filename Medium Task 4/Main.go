package main

import "fmt"

func main() {
	library := map[string][]string{}
	library["Alice"] = []string{"Book", "Magazine"}
	library["Bob"] = []string{"Dictionary"}
	library["Charlie"] = []string{"Newspaper"}
	library["David"] = []string{}
	countPeopleWithBooks := 0
	for _, value := range library {
		if len(value) > 0 {
			countPeopleWithBooks++
		}
	}
	fmt.Println("Количество читателей с изданиями на руках: ", countPeopleWithBooks)
	for key, value := range library {
		fmt.Println("У", key, " - ", len(value), "издания.")
	}
}
