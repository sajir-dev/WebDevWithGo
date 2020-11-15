package main

import "fmt"

func main() {

	// Challenge #1
	c := make(chan int)
	// c <- 1 // Created the deadlock
	go func() { // Solution => created a goroutine that reads from the channel
		c <- 1
	}()
	fmt.Println(<-c)

	// Challenge #2
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
