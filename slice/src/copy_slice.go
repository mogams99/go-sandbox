package src

import "fmt"

func CopySlice() {
	dst := make([]string, 3)
	src := []string{"dog", "cat", "rabbit", "hamster"}
	n := copy(dst, src)

	fmt.Println("dst:", dst)
	fmt.Println("src:", src)
	fmt.Println("n:", n)
}
