package src

import "fmt"

func BaseSlice() {
	fruits := []string{"apple", "banana", "grape", "orange"}

	aFruits := fruits[0:3]
	bFruits := fruits[1:4]

	aaFruits := aFruits[1:2]
	baFruits := bFruits[0:1]

	fmt.Printf("===============================================\n")
	fmt.Printf("===============================================\n")

	fmt.Println("Fruits:", fruits)
	fmt.Println("aFruits:", aFruits)
	fmt.Println("bFruits:", bFruits)
	fmt.Println("aaFruits:", aaFruits)
	fmt.Println("baFruits:", baFruits)

	fmt.Printf("===============================================\n")
	fmt.Printf("Change 'banana' to 'kiwi'\n")
	fmt.Printf("===============================================\n")
	baFruits[0] = "kiwi"

	fmt.Println("Fruits:", fruits)
	fmt.Println("aFruits:", aFruits)
	fmt.Println("bFruits:", bFruits)
	fmt.Println("aaFruits:", aaFruits)
	fmt.Println("baFruits:", baFruits)

	fmt.Printf("===============================================\n")
	fmt.Printf("Checking the capacity of slices\n")
	fmt.Printf("===============================================\n")
	fmt.Printf("Fruits: %v, len: %d, cap: %d\n", fruits, len(fruits), cap(fruits))
	fmt.Printf("aFruits: %v, len: %d, cap: %d\n", aFruits, len(aFruits), cap(aFruits))
	fmt.Printf("bFruits: %v, len: %d, cap: %d\n", bFruits, len(bFruits), cap(bFruits))
	fmt.Printf("aaFruits: %v, len: %d, cap: %d\n", aaFruits, len(aaFruits), cap(aaFruits))
	fmt.Printf("baFruits: %v, len: %d, cap: %d\n", baFruits, len(baFruits), cap(baFruits))

	fmt.Printf("===============================================\n")
	fmt.Printf("===============================================\n")

}
