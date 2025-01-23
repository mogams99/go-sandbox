package src

import "fmt"

func BaseMap() {
	data := map[string]int{}

	data["january"] = 50
	data["february"] = 40

	fmt.Println("January:", data["january"])
	fmt.Println("February:", data["february"])
	fmt.Println("March:", data["March"])

}
