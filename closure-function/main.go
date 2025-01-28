package main

import "fmt"

func main() {
	var findMinMax = func(numbers []int) (int, int) {
		var min, max int
		for _, number := range numbers {
			if number > max {
				max = number
			}
			if number < min {
				min = number
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = findMinMax(numbers)

	fmt.Printf("Data: %v\n", numbers)
	fmt.Printf("Min: %v\n", min)
	fmt.Printf("Max: %v\n", max)
}
