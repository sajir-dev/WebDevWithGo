package main

import "fmt"

func isSubsequence(sub string, big string) bool {
	s := []byte(sub)
	ss := []byte(big)

	i := 0
	for j := 0; j < len(ss); j++ {
		if s[i] == ss[j] {
			i++
		}
		if len(s) == i { // since i++ makes the value of i index+1 at 12th line
			return true
		}
	}
	return false
}

func main() {
	a := "abcabc"
	b := "aabcdfbsabdc"
	fmt.Println(isSubsequence(a, b))
}
