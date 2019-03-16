package main

import "fmt"

func main() {

	s := "你好啊，世界！"

	for i, ch := range s {
		fmt.Printf("(%d, %c)", i, ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}

}
