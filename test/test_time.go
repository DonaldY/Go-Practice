package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	time1 := "2019-03-20 08:50:29"
	time2 := "2019-03-21 09:04:25"

	t1, _ := time.Parse("2006-01-02 15:04:05", time1)
	t2, _ := time.Parse("2006-01-02 15:04:05", time2)

	fmt.Println(now)

	fmt.Println(t1.After(t2))
	fmt.Println(now.After(t1))

}
