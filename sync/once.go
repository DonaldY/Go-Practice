package main


import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Once是sync包中的一个对象，它只有一个方法Do。
在程序运行过程中，无论被多少次调用，只会执行一次，就与结构体的名称一样，once（一次）。
 */
func main() {
	var once sync.Once

	for i := 0 ; i <= 10; i++ {

		go once.Do(func() {

			fmt.Println("hello world")

		})
	}

	time.Sleep(time.Second * 2)
}


/**
type Once struct {
	m    Mutex
	done uint32
}

ｍ是了保证并发安全性的，
done是标志是否已经执行过此方法，如果done是１则表示执行过，０表示未执行。
 */