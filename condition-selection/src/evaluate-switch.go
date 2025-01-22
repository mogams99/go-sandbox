package src

import "fmt"

func EvaluateSwitch(param int) {
	switch param {
	case 10:
		fmt.Println("evaluateSwitch: param is 10")
	case 5, 6, 7, 8, 9:
		fmt.Println("evaluateSwitch: param is between 5 and 9")
	default:
		fmt.Println("evaluateSwitch: param is less than 5")
	}
}
