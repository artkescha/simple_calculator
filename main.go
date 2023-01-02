package main

import "fmt"

func main() {
	var wantResult int
	fmt.Println("Enter want result:")
	fmt.Scanf("%d\n", &wantResult)
	fmt.Printf("iteration metod stepsCount: %d\n", iteration(wantResult))
	fmt.Printf("recursive metod stepsCount: %d\n", recursive(wantResult))
}
