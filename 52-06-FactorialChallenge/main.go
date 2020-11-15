package main

import "fmt"

func main() {
	// c := factorial(4)

	in := gen()

	f := factorial(in)

	for n := range f {
		// fmt.Println(<-c)
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(num <-chan int) <-chan int {
	// var c chan int
	out := make(chan int)
	go func() {
		for n := range num {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
