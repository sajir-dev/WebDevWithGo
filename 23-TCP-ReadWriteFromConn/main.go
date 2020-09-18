package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// listening to a port
	li, err := net.Listen("tcp", ":3080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Accepting a connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// launching a seperate goroutine to handle a connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s, \n", ln)
	}
	defer conn.Close()

	fmt.Println("You are here")
}
