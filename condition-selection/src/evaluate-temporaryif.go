package src

import "fmt"

func EvaluateTemporaryIf(param float32) {
	if percent := param / 100; percent >= 100 {
		fmt.Printf("evaluateTemporaryIf: percent is greater than 100 | %.2f%%\n", percent)
	} else if percent >= 50 {
		fmt.Printf("evaluateTemporaryIf: percent is greater than 50 | %.2f%%\n", percent)
	} else {
		fmt.Printf("evaluateTemporaryIf: percent is less than 50 | %.2f%%\n", percent)
	}
}
