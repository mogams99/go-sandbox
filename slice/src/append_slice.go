package src

import "fmt"

func AppendSlice() {
	animals := []string{"dog", "cat", "rabbit"}
	aAnimals := append(animals, "hamster")

	fmt.Println("Animals:", animals)
	fmt.Println("aAnimals:", aAnimals)

	fmt.Printf("===============================================\n")
	fmt.Printf("If len of slice is equal to cap of slice\n")
	fmt.Printf("Then, new element is new reference\n")
	fmt.Printf("But if len of slice is less than cap of slice\n")
	fmt.Printf("Then, new element is placed into the capacity range, making all other slice elements with the same reference change their values\n")
	fmt.Printf("===============================================\n")

	bAnimals := animals[0:2]
	fmt.Printf("len: %d, cap: %d\n", len(bAnimals), cap(bAnimals))

	fmt.Println("Animals:", animals)
	fmt.Println("bAnimals:", bAnimals)

	fmt.Printf("===============================================\n")
	fmt.Printf("Append 'hamster' to bAnimals\n")
	fmt.Printf("===============================================\n")
	cAnimals := append(bAnimals, "hamster")

	fmt.Println("Animals:", animals)
	fmt.Println("bAnimals:", bAnimals)
	fmt.Println("cAnimals:", cAnimals)

	fmt.Printf("===============================================\n")
	fmt.Printf("===============================================\n")

}
