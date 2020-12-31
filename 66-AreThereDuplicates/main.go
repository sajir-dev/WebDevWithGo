package main

import "fmt"

func areThereDuplicates(s ...string) bool {
	// areThereDuplicates("a", "b", "c", "a") // true
	// areThereDuplicates(1, 5, 8, 9) // false

	var b []byte

	for _, v := range s {
		u := []byte(string(v))
		b = append(b, u...)
		fmt.Println(v, u, b)
	}

	// bs := []byte(s)

	m := make(map[byte]int)

	for _, v := range b {
		m[v]++
		if m[v] > 1 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(areThereDuplicates("a", "2", "3", "8", "9", "a"))
}
