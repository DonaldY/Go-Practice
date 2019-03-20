package main

import "fmt"

// 数组的指针：*[3]int
// 指针数组：[2]*int

type GameRecordVo struct {
	Phone  string `json:"phone"`
	Amount int    `json:"amount"`
}

func main() {

	gameRecordVos := [5]GameRecordVo{{"13333333333", 112}, {"13333333334", 113},
		{"13333333335", 114}, {"13333333336", 115}, {"13333333337", 116}}

	fmt.Println(gameRecordVos)

	// _ 意味忽略，下面例子能成功赋值
	for i, vo := range gameRecordVos  {

		phone := vo.Phone

		pre := phone[:3]
		last := phone[7:]

		gameRecordVos[i].Amount = 123
		gameRecordVos[i].Phone = pre + "****" + last
	}

	fmt.Println(gameRecordVos)
}
