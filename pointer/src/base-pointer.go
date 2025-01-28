package src

import "fmt"

func BasePointer() {
	fmt.Println("==========================================")
	fmt.Println("\t \t Pointer")
	fmt.Println("==========================================")

	var numberA int = 10
	var numberB *int = &numberA

	fmt.Println("numberA (value) \t :", numberA)
	fmt.Println("numberA (address) \t :", &numberA)

	fmt.Println("numberB (value) \t :", *numberB)
	fmt.Println("numberB (address) \t :", numberB)

	fmt.Println("==========================================")
	fmt.Println("\t Change Pointer Value")
	fmt.Println("==========================================")

	numberA = 15

	fmt.Println("numberA (value) \t :", numberA)
	fmt.Println("numberA (address) \t :", &numberA)

	fmt.Println("numberB (value) \t :", *numberB)
	fmt.Println("numberB (address) \t :", numberB)

	fmt.Println("==========================================")
}
