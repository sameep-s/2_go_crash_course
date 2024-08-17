package main

import (
	"errors"
	"fmt"
)

func main() {
	printVal := "Print Value."
	printMe(printVal)

	var remainder, result, err = division(4, 23)
	if err != nil {
		println(err)
	} else if remainder == 0 {
		fmt.Printf("The division value is %v with remainder zero\n", result)
	} else {
		fmt.Printf("The division result is %v with remainder %v\n", result, remainder)
	}

}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func division(numerator int, denominator int) (int, int, error) {

	if denominator == 0 {
		err := errors.New("denominator cannot be zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return remainder, result, nil

}
