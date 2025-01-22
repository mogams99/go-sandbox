package main

import "fmt"

func main() {
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("Left && Right: \t(%t)\n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("Left || Right: \t(%t)\n", leftOrRight)

	notLeft := !left
	fmt.Printf("!Left: \t\t(%t)\n", notLeft)
}
