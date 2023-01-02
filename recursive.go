package main

func recursive(wantResult int) int {
	currentValue := 0
	stepsCount := 0
	_, stepsCount = recCalculate(wantResult, currentValue, stepsCount)
	return stepsCount
}

func recCalculate(wantResult, currentValue, stepsCount int) (int, int) {
	if wantResult <= currentValue {
		return currentValue, stepsCount
	}
	currentValue1 := currentValue + 1
	currentValue2 := currentValue * currentValue

	currentValue = max(currentValue1, currentValue2)
	stepsCount++
	return recCalculate(wantResult, currentValue, stepsCount)
}
