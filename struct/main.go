package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var bachelor student
	bachelor.name = "Mochammad Gamal"
	bachelor.grade = 2

	fmt.Println("Name \t :", bachelor.name)
	fmt.Println("Grade \t :", bachelor.grade)
}
