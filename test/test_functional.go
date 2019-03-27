package main

import "fmt"

type test struct {
	Num uint
}

func main() {

	u := Func()

	fmt.Println(u)
}

func Func() (ret test) {

	ret = test{Num: 1}

	return ret
}
