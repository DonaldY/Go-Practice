package main

import (
	"fmt"
	"sync"
)

// 对象内部有一个计数器，最初从0开始
var wg sync.WaitGroup

func printFunc() {
	fmt.Println("Im a goroutine")
	wg.Done()
}

// goroutine1
func main() {

	defer fmt.Println("---main goroutine over---")

	wg.Add(1)

	// goroutine2
	go printFunc()

	fmt.Println("---wait sub goroutine over---")

	wg.Wait()

	fmt.Println("---sub goroutine over---")
}
