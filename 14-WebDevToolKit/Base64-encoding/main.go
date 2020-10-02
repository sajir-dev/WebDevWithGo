package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "der text commonly used in the graphic, print, and publishing industries for previewing layouts and visual moc"
	fmt.Println(s)
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("decoding failed", err)
	}
	fmt.Println(string(bs))
}
