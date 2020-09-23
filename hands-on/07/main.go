package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// listen to port 80; listen is a function of net/http package and listen is attahced with net struct in the package. Here we should pass the listen type tcp as first parameter and the specify the port number
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("unable to connect to the port 8080:", err)
	}
	defer li.Close()

	for {
		// Accept the connections from the line established with port 8080 with an infinite loop
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("END OF THE END OF HTTP REQUEST HEADERS")
			break
		}
	}
}
