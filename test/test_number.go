package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int = 100

	var un uint = uint(n)

	un2, _ := strconv.Atoi("100")

	var un3 uint = uint(un2)

	fmt.Println(un)

	fmt.Println(un2)

	fmt.Println(un3)
}
