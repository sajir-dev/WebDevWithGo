package main

import (
	"fmt"
	"sync"
)

func main() {
	c := gen()

	c0 := factorial(c)
	c1 := factorial(c)
	c2 := factorial(c)
	c3 := factorial(c)
	c4 := factorial(c)
	c5 := factorial(c)
	c6 := factorial(c)
	c7 := factorial(c)
	c8 := factorial(c)
	c9 := factorial(c)

	var num int = 0

	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		num++
		fmt.Println(num, "\t", n)
	}
}

func gen() <-chan int {
	num := make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			for j := 3; j < 13; j++ {
				num <- j
			}
		}
		close(num)
	}()
	// close(num)
	return num
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	var f int = 1
	for i := n; i > 0; i-- {
		f *= i
	}
	return f
}

func merge(cs ...<-chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, v := range cs {
		go output(v)
	}

	// Start a goroutine to close out once all the output goroutines are done. This must start after the wg.Add call
	// This must start after wg.Add() call
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
