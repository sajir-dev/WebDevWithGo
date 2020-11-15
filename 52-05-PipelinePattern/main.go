package main

import "fmt"

func main() {
	// Setup the pipeline and consume the output
	for n := range sq(sq(gen(4, 5))) {
		fmt.Println(n)
	}
}

func gen(num ...int) chan int {
	c := make(chan int)
	go func() {
		for _, n := range num {
			c <- n
		}
		close(c)
	}()
	return c
}

func sq(in chan int) chan int {
	c := make(chan int)
	go func() {
		for n := range in {
			c <- n * n
		}
		close(c)
	}()
	return c
}
