package main

import "fmt"

// #1 Normal programming to get factorial
// func main() {
// 	fact := factorial(4)
// 	fmt.Println(fact)
// }

// func factorial(num int) int {
// 	f := 1
// 	go func() {
// 		for i := num; i > 0; i-- {
// 			f *= i
// 		}
// 	}()
// 	return f
// }

// #2 Concurrent programming to get factorial
func main() {
	c := factorial(4)
	for n := range c {
		// fmt.Println(<-c)
		fmt.Println(n)
	}
}

func factorial(num int) chan int {
	// var c chan int
	c := make(chan int)
	go func() {
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
