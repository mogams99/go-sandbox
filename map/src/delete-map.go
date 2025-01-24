package src

import "fmt"

func DeleteMap() {
	data := map[string]int{
		"january":  50,
		"february": 40,
	}

	fmt.Println(len(data))
	fmt.Println(data)

	delete(data, "january")

	fmt.Println(len(data))
	fmt.Println(data)
}
