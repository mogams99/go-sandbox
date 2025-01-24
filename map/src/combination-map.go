package src

import "fmt"

func CombinationMap() {
	data := []map[string]string{
		{"name": "wick", "grade": "A"},
		{"name": "jason", "grade": "B"},
		{"name": "ethan", "grade": "C"},
	}

	fmt.Printf("Data: %v\n", data)

	for _, value := range data {
		fmt.Printf("Name: %s, Grade: %s\n", value["name"], value["grade"])
	}
}
