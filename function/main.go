package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	names := []string{"Mochammad", "Gamal"}
	printMessage("Good morning,", names)

	fmt.Println("========================================")
	fmt.Println("Run Random With Range Function")
	fmt.Println("========================================")

	var count int
	for {
		randomValue := randomWithRange(2, 10)
		fmt.Println("Random Number:", randomValue)
		count++
		if count >= 5 {
			break
		}
	}
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	value := randomizer.Int()%(max-min+1) + min
	return value
}
