package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]

	fmt.Printf("%v, %T, %v\n", indexed, indexed, myString)

	for i, v := range myString {
		fmt.Printf("Index: %v, val: %v.\n", i, v)
	}

	var strBuilder strings.Builder
	var strSlice = []string{"h", "e", "l", "l", "o"}

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var castStr = strBuilder.String()
	fmt.Println("CastStr: ", castStr)

}
