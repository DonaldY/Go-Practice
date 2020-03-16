package main

import (
	"fmt"
	"sync"
)

/**
  go 关键字创建的 goroutine 内新增了 for 无限循环打印输出，子线程会随主线程销毁而销毁
 */

// 对象内部有一个计数器，最初从0开始
var wg sync.WaitGroup

// goroutine1
func main() {

	defer fmt.Println("---main goroutine over---")

	wg.Add(1)

	// goroutine2
	go func() {
		fmt.Println("Im a goroutine")
		wg.Done()

		// 无限循环
		for {
			fmt.Println("---sub goroutine---")
		}
	}()

	fmt.Println("---wait sub goroutine over---")

	wg.Wait()

	fmt.Println("---sub goroutine over---")
}
