package main

import "sync"

func TestSync() {

	var wg sync.WaitGroup

	for i := 1; i < 100; i ++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// ...
		}(i)
	}
}
