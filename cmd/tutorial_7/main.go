package main

import (
	"fmt"
)

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value stored at p is %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)

	*p = 23
	fmt.Printf("The value of p is %v\n", *p)
	fmt.Printf("The value of &p is %v\n", &p) // dereferencing the pointer
}
