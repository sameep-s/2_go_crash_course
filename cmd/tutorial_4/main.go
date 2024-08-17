package main

import (
	"fmt"
)

func main() {

	// Simple array initilization
	var intArr [3]int32
	fmt.Println(intArr[1:3])

	// Slice - wrapper around arrays
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(append(intSlice, 44))

	// Using slice using make function
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length and capacity of the slice is len: %v, cap: %v\n", len(intSlice3), cap(intSlice3))

	// Create a slice using spread operator
	var intSlice2 []int32 = append(intSlice, intSlice3...)
	fmt.Println(intSlice2)

	// ------------------------------------------------------------- MAP

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 22}
	fmt.Println(myMap2)

	// delete(myMap2, "Adam") // deleting a value
	var age, ok = myMap2["Adam"]

	if ok {
		fmt.Printf("age %v\n", age)
	} else {
		fmt.Printf("The key doesn't exist\n")
	}

	// ------------------------------------------------------------- Loop
	for name := range myMap2 {
		fmt.Printf("Name %v\n", name)
	}

	// While loop in go
	var i int32 = 10

	for i > 0 {
		fmt.Println(i)
		i--
	}

}
