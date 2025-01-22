package main

import "condition-selection/src"

func main() {
	const (
		number  = 8
		decimal = 8840.0
	)

	src.EvaluateKeywordIf(number)
	src.EvaluateTemporaryIf(decimal)
	src.EvaluateSwitch(number)
}
