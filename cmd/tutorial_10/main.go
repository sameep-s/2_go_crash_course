package main

import (
	"fmt"
)

// Generics in GO Programming
func main() {
	var intSlice = []int{1, 2, 3, 22, 33, 44, 44, 55, 65, 32}
	var sum int = sumSlice(intSlice)

	fmt.Printf("Sum: %v\n", sum)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, val := range slice {
		sum += val
	}
	return sum
}
