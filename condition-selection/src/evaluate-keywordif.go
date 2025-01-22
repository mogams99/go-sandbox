package src

import "fmt"

func EvaluateKeywordIf(param int) {
	if param == 10 {
		fmt.Println("EvaluateKeywordIf: param is 10")
	} else if param > 5 {
		fmt.Println("EvaluateKeywordIf: param is greater than 5")
	} else if param == 4 {
		fmt.Println("EvaluateKeywordIf: param is 4")
	} else {
		fmt.Println("EvaluateKeywordIf: param is less than 4")
	}
}
