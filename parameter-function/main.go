package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"one", "two", "three", "four", "five"}

	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	dataLengthMoreThanEqualFive := filter(data, func(each string) bool {
		return len(each) >= 5
	})

	fmt.Println("Data: ", data)
	fmt.Println("Data contains 'o': ", dataContainsO)
	fmt.Println("Data length more than equal 5: ", dataLengthMoreThanEqualFive)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
