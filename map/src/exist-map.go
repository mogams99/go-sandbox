package src

import "fmt"

func ExistMap() {
	data := map[string]int{
		"january":  50,
		"february": 40,
	}

	value, exist := data["mei"]

	if exist {
		fmt.Println("Value exist: ", value)
	} else {
		fmt.Println("Value doesn't exist")
	}
}
