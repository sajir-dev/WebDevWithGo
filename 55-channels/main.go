package main

import (
	"fmt"
	"runtime"
)

// func main() {
// 	var c chan int
// 	go func() {
// 		for i := 3; i < 10; i++ {
// 			c <- fact(i)
// 		}
// 		close(c)
// 	}()
// 	fmt.Println(<-c)
// }

// func fact(n int) int {
// 	var acc int
// 	for i := n; i > 0; i++ {
// 		acc = acc * n
// 	}
// 	return acc
// }

// func main() {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	c := fact(5)
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// func fact(lim int) chan int {
// 	c := make(chan int)
// 	go func() {
// 		acc := 1
// 		for i := lim; i > 0; i-- {
// 			acc *= i
// 		}
// 		c <- acc
// 		close(c)
// 	}()
// 	return c
// }

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("In goroutine:", i)
			c <- i
		}
		close(c)
		fmt.Println("1.", runtime.NumCPU())
		fmt.Println("2.", runtime.NumGoroutine())
	}()

	for v := range c {
		fmt.Println("3.", runtime.NumCPU())
		fmt.Println("4.", runtime.NumGoroutine())
		fmt.Println("In main:", v)
	}
}
