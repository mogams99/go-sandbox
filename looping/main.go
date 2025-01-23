package main

import "looping/src"

func main() {
	const (
		count      = 1
		countLimit = 10
	)

	src.ForLoop(count, countLimit)
	src.ConditionLoop(count, countLimit)
	src.PlainLoop(count, countLimit)
	src.RangeLoop(count, countLimit)
}
