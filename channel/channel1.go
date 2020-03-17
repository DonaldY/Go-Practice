package main

import (
	"fmt"
	"time"
)

/**
channel 理解为并发安全的队列
 */

func main() {

	// 无缓冲通道 quit 来做反向通知子线程停止
	quit := make(chan struct{})

	go func() {

		for {
			select {
			case <-quit:
				fmt.Println("sub goroutine is over")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("sub goroutine do something")
			}
		}
	}()

	time.Sleep(time.Second * 3)

	fmt.Println("main goroutine start stop sub goroutine")

	// 关闭通道 quit
	close(quit)

	time.Sleep(time.Second * 10)
	fmt.Println("main goroutine is over")
}
