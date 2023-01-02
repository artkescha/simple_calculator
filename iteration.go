package main

func iteration(wantResult int) int {
	stepsCount := 0
	currentValue := 0
	for currentValue < wantResult {
		currentValue = calculate(currentValue)
		stepsCount++
	}
	return stepsCount
}

func calculate(currentValue int) int {
	return max(currentValue+1, currentValue*currentValue)
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
