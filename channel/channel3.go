package main

import "fmt"

/**
生产者/消费者模型
 */

func printer(ch <- chan int, wg chan <- int) {

	for i := range ch{

		fmt.Println(i)
	}

	wg <- 1

	close(wg)
}

func main() {

	// 1. 创建缓冲通道
	ch := make(chan int, 10)

	// 2. 创建同步用的无缓冲通道
	wg := make(chan int)

	// 3. 开启 goroutine
	go printer(ch, wg)

	// 4. 写入到通道
	for i := 1; i < 100 ; i ++ {
		ch <- 1
	}

	// 5. 关闭通道
	close(ch)

	fmt.Println("wait sub goroutine over")

	// 6. 等待子goroutine结束
	<- wg
	fmt.Println("main goroutine over")
}
