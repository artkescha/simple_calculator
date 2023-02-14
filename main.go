package main

import "fmt"

func main() {
	var wantResult int
	fmt.Println("Enter want result:")
	fmt.Scanf("%d\n", &wantResult)
	fmt.Printf("iteration metod stepsCount: %d\n", iteration(wantResult))
	_, stepsCount := recursive(wantResult, 0, 0)
	fmt.Printf("recursive metod stepsCount: %d\n", stepsCount)
}
