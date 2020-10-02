package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c1 := getCode("test@email.com")
	fmt.Println(c1)
	c2 := getCode("test@email.com")
	fmt.Println(c2)
	fmt.Println(c1 == c2)
}

func getCode(c string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, c)
	return fmt.Sprintf("%x", h.Sum(nil))
}
