package src

import "fmt"

func ForRangeMap() {
	data := map[string]int{
		"january":  50,
		"february": 40,
		"march":    30,
		"april":    20,
		"may":      10,
	}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
