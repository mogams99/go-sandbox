package src

import "fmt"

func VerticalArray(size int64) {
	fmt.Print("Array size is ", size, "\n")

	var fruits = make([]string, size)
	fruits = []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
	}

	fruits = fruits[:size]
	fmt.Println(fruits)

}
