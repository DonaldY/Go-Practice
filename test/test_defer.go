package main

import "fmt"

func main() {

	fmt.Println(*testDefer())

}

func testDefer() *int {

	a := 1

	defer func() {
		fmt.Println("123456")
		a = 3
	}()

	a = 2

	return &a
}
