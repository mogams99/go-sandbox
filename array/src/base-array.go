package src

import "fmt"

func BaseArray(size int) {
	// The above line will not work because make() expects a type, not a string.
	// The following line will work because it uses the typeName argument to create the array.
	names := make([]string, size)
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"
	names[4] = "Pete"

	fmt.Println("The length of the array is", len(names))
	fmt.Println(names[0], names[1], names[2], names[3], names[4])

}
