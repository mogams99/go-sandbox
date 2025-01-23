package src

import "fmt"

func ForLoopArray() {
	animals := [...]string{"dog", "cat", "bird", "fish", "lion"}

	for i := 0; i < len(animals); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, animals[i])
	}
}
