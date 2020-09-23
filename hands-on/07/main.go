package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("unable to listen to 8080:", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			go serve(conn, scanner)
		}
	}
}

func serve(conn net.Conn, scanner *bufio.Scanner) {
	ln := scanner.Text()
	fmt.Println(ln)
}
