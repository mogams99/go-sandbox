package src

import "fmt"

func ConditionLoop(count int, countLimit int) {
	for count <= countLimit {
		fmt.Println("Condition loop:", count)
		count++
	}
}
