package main

import "fmt"

func main() {
	s1 := "abracadabra"
	s2 := "dabra"
	ss := []byte(s1)
	s := []byte(s2)
	fmt.Println(naiveSearch(ss, s))
}

func naiveSearch(ss []byte, s []byte) bool {
	for i := 0; i < len(ss); i++ {
		// fmt.Println("#1", i, j, len(ss), len(s))
		for j := 0; j < len(s) && i < len(ss); j++ {
			if s[j] != ss[i+j] {
				break
			}
			// fmt.Println("#2", i, j)
			i++
			j++
			if j == len(s) {
				// fmt.Println("#3", i, j)
				return true
			}
		}
	}
	return false
}
