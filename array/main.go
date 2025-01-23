package main

import (
	"array/src"
)

func main() {
	const (
		arraySize = 5
	)

	src.BaseArray(arraySize)
	src.VerticalArray(arraySize)
	src.NoSizeArray()
	src.MultiDimensionalArray()
	src.ForLoopArray()
	src.ForRangeArray()
}
