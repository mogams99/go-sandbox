package src

import "fmt"

func NoSizeArray() {
	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Printf("The length of the array is %d\n", len(numbers))
	fmt.Printf("Data array: %v\n", numbers)
}
