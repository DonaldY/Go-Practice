package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(str string) int {

	lastOccurred := make(map[int32]int)
	start := 0
	maxLength := 0

	for i, ch := range str {

		// 如果在这下标之前存在，则过
		// 如果在这下标之后存在，
		if lastI, ok := lastOccurred[ch]; ok && lastI >= lastOccurred[ch] {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abc"),
		lengthOfNonRepeatingSubStr(""),
		lengthOfNonRepeatingSubStr("abcabcaaa"),
		lengthOfNonRepeatingSubStr("一二三"))
}
