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
	j := 0
	for i := 0; i < len(ss); i++ {
		// fmt.Println("#1", i, j, len(ss), len(s))
		var match int
		if ss[i] == s[j] {
			match = i
			for j < len(s) && i < len(ss) {
				if s[j] != ss[i] {
					i = match
					j = 0
					break
				}
				fmt.Println("#2", i, j)
				i++
				j++
				if j == len(s) {
					fmt.Println("#3", i, j)
					return true
				}
			}
		}
	}
	return false
}
