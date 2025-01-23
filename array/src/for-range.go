package src

func ForRangeArray() {
	plants := [...]string{"apple", "banana", "cherry", "date", "elderberry"}

	for i, plant := range plants {
		println("Index: ", i, ", Value: ", plant)
	}
}
