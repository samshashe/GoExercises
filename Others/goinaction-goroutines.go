package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main1() {

	//runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go doWork("A")
	go doWork("B")

	// Give the goroutines time to run.
	time.Sleep(3 * time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")

	// Wait for the goroutines to finish.
	wg.Wait()
}

func doWork(name string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 5; count++ {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

	}
}
