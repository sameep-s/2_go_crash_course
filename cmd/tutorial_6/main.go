package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func main() {

	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 2}
	myEngine.mpg = 30
	myEngine.owner.name = "Sameep"

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner.name)

}
