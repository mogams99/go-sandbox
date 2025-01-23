package src

import "fmt"

func MultiDimensionalArray() {
	numbers_first := [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	numbers_second := [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("Number fist array: ", numbers_first)
	fmt.Println("Number second array: ", numbers_second)
}
