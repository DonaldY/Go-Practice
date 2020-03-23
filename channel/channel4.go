package main

func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, n := range  nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func main() {


}
