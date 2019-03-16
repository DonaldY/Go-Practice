package main

import (
	"fmt"
)

var name int

func grade(score int) string {

	g := ""
	switch {
	case score < 0 || score > 100:
		fmt.Printf("Wrong score: %d", score)
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func switch_test(a int) {

	switch a {
	case 1:
	case 2: fallthrough
	default:
	}
}

func getResult() (result int, err error) {
	return 1, nil
}

func swap(a, b *int) {

	*b, *a = *a, *b
}

func main()  {

	fmt.Println(name)

	fmt.Println(grade(100), grade(99))

	if result, err := getResult(); err == nil {
		fmt.Println(result)
	}

	a, b := 1, 2
	swap(&a, &b)
	m := map[string]string {
		"haha": "lala",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}


