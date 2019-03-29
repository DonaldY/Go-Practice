package main

import "fmt"

func main() {

	testDefer()

}

func testDefer() {

	a := 1

	defer func() {
		fmt.Println(a)
	}()

	a = 2
}
