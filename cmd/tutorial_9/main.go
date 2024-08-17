package main

import (
	"fmt"
	"time"
)

func main() {
	var t0 = time.Now()
	var c = make(chan int, 10)
	go process(c)

	for i := range c {
		fmt.Printf("The val in channel is: %v\n", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("Time elapsed %v\n", time.Since(t0))
}

func process(c chan int) {
	defer close(c)

	for i := 0; i <= 9; i++ {
		fmt.Printf("Adding val: %v\n", i+1)
		c <- i + 1
	}
	fmt.Println("Exiting process")
}
