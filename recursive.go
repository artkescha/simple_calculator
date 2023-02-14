package main

func recursive(wantResult, currentValue, stepsCount int) (int, int) {
	if wantResult <= currentValue {
		return currentValue, stepsCount
	}

	currentValue = max(currentValue+1, currentValue*currentValue)
	stepsCount++
	return recursive(wantResult, currentValue, stepsCount)
}
