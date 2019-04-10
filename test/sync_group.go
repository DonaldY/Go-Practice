package main

import (
	"fmt"
	"sync"
)

func main() {

	TestSync()
}

func TestSync() {

	var wg sync.WaitGroup

	for i := 1; i < 100; i ++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finished!")
}
