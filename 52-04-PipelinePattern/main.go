package main

import "fmt"

func main() {
	// setup a pipeline
	c := gen(2, 3, 4)
	out := sq(c)

	// Consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
