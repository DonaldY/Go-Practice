package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

func main() {
	test1()
	test2()
}

// count++  并发不安全
func test1()  {
	var wg sync.WaitGroup
	count := 0
	t := time.Now()
	for i := 0 ; i < 100000 ; i++ {
		wg.Add(1)
		go func(i int) {
			count++   //count不是并发安全的
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test1 花费时间：%d, count的值为：%d \n",time.Now().Sub(t),count)
}

// atomic.AddInt64(&count,1)  //原子操作
func test2()  {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0 ; i < 100000 ; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count,1)  //原子操作
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test2 花费时间：%d, count的值为：%d \n",time.Now().Sub(t),count)
}