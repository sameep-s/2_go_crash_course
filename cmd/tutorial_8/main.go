package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"apple", "banana", "cherry", "date"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := range dbData {
		go dbCall(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Println("The results are: ", results)
}

func dbCall(i int) {
	// Simulate DB call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(str string) {
	m.Lock()
	results = append(results, str)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The value from db call is %v\n", results)
	m.RUnlock()
}
