package main

import "fmt"

func main() {
	avg := calculate(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	message := fmt.Sprintf("The average is %.2f", avg)
	fmt.Println(message)
}

func calculate(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg
}
