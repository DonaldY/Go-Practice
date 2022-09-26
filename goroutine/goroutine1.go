package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// goroutine1
func main() {

	// defer用于资源的释放，会在函数返回之前进行调用
	defer fmt.Println("====主goroutine 结束====")

	wg.Add(1)

	// 运行匿名函数
	go func() {
		fmt.Println("运行子goroutine")
		wg.Done()
	}()

	fmt.Println("====等待 子goroutine 结束====")

	wg.Wait()

	fmt.Println("====子goroutine 结束====")
}
