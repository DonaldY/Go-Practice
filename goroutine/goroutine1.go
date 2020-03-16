package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// goroutine1
func main() {

	defer fmt.Println("---main goroutine over---")

	wg.Add(1)

	// goroutine2
	go func() {
		fmt.Println("Im a goroutine")
		wg.Done()
	}()

	fmt.Println("---wait sub goroutine over---")

	wg.Wait()

	fmt.Println("---sub goroutine over---")
}
