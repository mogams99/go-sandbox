package main

import "fmt"

func main() {
	value := (((2+6)%3)*4 - 2) / 3
	counter := 2
	isEqual := (value == counter)
	isNotEqual := (value != counter)
	isLessThan := (value < counter)
	isLessThanOrEqual := (value <= counter)
	isGreaterThan := (value > counter)
	isGreaterThanOrEqual := (value >= counter)
	fmt.Println("========================================")
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Counter: %d\n", counter)
	fmt.Println("========================================")
	fmt.Printf("Value == Counter: %t\n", isEqual)
	fmt.Printf("Value != Counter: %t\n", isNotEqual)
	fmt.Printf("Value < Counter: %t\n", isLessThan)
	fmt.Printf("Value <= Counter: %t\n", isLessThanOrEqual)
	fmt.Printf("Value > Counter: %t\n", isGreaterThan)
	fmt.Printf("Value >= Counter: %t\n", isGreaterThanOrEqual)
	fmt.Println("========================================")
}
