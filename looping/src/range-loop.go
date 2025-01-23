package src

import "fmt"

func RangeLoop(count int, countLimit int) {
	for count, limit := range []int{count, countLimit} {
		fmt.Println("Range loop | Index=", count, "Value=", limit)
	}
}
