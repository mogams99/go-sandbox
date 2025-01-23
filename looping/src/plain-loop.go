package src

import "fmt"

func PlainLoop(count int, countLimit int) {
	for {
		fmt.Println("Plain loop:", count)
		count++

		if count > countLimit {
			break
		}
	}
}
