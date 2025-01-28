package src

import "fmt"

func ParameterPointer() {
	fmt.Println("******************************************")
	fmt.Println("==========================================")
	fmt.Println("\t Parameter Pointer")
	fmt.Println("==========================================")

	number := 7
	fmt.Println("Before \t :", number)

	change(&number, 10)
	fmt.Println("After \t :", number)
	fmt.Println("==========================================")
}

func change(ori *int, value int) {
	*ori = value
}
