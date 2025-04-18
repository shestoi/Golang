package main

import "fmt"

func contains(a []string, x string) bool {
	for _, str := range a {
		if x == str {
			return true
		}
	}
	return false
}
func getMax(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	fmt.Println(getMax(1, 2, 3, 4, 5, 6, 7))
	var Names = []string{"Ahmed", "Magomed", "Usman", "Timur"}
	fmt.Println(contains(Names, "Timur"))

}
