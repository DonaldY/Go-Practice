package main

import (
	"fmt"
	"sort"
)

var c chan int

var nums [] int

func main() {

	// 1. 初始化通道
	c = make(chan int)

	nums = []int{5, 3, 7, 2, 9, 1, 6}

	// 2. 开启 goroutine 排序
	go func() {

		sort.Ints(nums)

		c <- 1
	}()

	// 3. 阻塞，直到通道内有元素
	<-c
	fmt.Println(nums)
}
