package main

import "fmt"

func main() {
	const firstName string = "john"
	fmt.Print("hello ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	const (
		shape          = "square"
		isToday  bool  = true
		numeric  uint8 = 102
		floatNum       = 3.14
	)
	fmt.Println("====================")
	fmt.Println(shape)
	fmt.Println(isToday)
	fmt.Println(floatNum)

	const (
		a = "constant a"
		b
	)
	fmt.Println("====================")
	fmt.Println(a)
	fmt.Println(b)

	const (
		today string = "sunday"
		now
		isToday2 = true
	)
	fmt.Println("====================")
	fmt.Println(today)
	fmt.Println(now)
	fmt.Println(isToday2)

	const one, two = 1, 2
	const three, four string = "three", "four"
	fmt.Println("====================")
	fmt.Println(one, two)
	fmt.Println(three, four)
}
