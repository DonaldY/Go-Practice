package main

import "fmt"

type Test struct {
	Name  string
	Age   int
}

// 探究对象赋值
func main() {

	test := Test{Name:"duoNa", Age:12}

	if (Test{}) != test {

		fmt.Println("test isn't nil")
	}

	testReference(test)
}

func testReference(test Test) {

	fmt.Println(test)

	str := "12321"

	age := 123

	test.Name = str

	test.Age = age

	fmt.Println(test)
}
