package main

import (
	"fmt"
	"sync"
	"time"
)

func printNums(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Hello World")

	var wg sync.WaitGroup
	wg.Add(1) // Increment the counter for the goroutine

	go printNums(&wg)

	wg.Wait() // Wait for all goroutines to complete
}
