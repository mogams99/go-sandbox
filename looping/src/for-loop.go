package src

import "fmt"

func ForLoop(count int, countLimit int) {
	for i := count; i <= countLimit; i++ {
		fmt.Println("For loop:", i)
	}
}
