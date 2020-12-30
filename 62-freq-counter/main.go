package main

import "fmt"

func freqCounter(n []int, nsq []int) bool {
	// [1,2,3] , [1, 9, 4] => true
	// [1, 1, 4] , [16, 16, 1] => false
	var objN = make(map[int]int)
	var objN2 = make(map[int]int)

	if len(n) != len(nsq) {
		return false
	}

	for _, v := range n {

		if objN[v] > 0 {
			objN[v]++
		} else {
			objN[v] = 1
		}
	}

	for _, v := range nsq {
		if objN2[v] > 0 {
			objN2[v]++
		} else {
			objN2[v] = 1
		}
	}

	for k := range objN {
		// fmt.Println(k, ":", objN[k])
		if objN[k] != objN2[k*k] {
			return false
		}
	}

	return true
}

// anagram
func anagram(a string, b string) bool {
	// abcd , acbd are anagrams
	// abcdee, eeabcd are anagrams
	// pqrst, abpqr are not anagrams

	c := []rune(a)
	d := []rune(b)

	if len(a) != len(b) {
		return false
	}

	obj1 := make(map[rune]int)
	obj2 := make(map[rune]int)

	for _, v := range c {
		obj1[v]++
	}

	for _, v := range d {
		obj2[v]++
	}

	for k := range obj1 {
		if obj1[k] != obj2[k] {
			return false
		}
	}

	return true
}

func main() {
	// a := []int{2, 4, 2}
	// b := []int{4, 16, 4}

	// fmt.Println(freqCounter(a, b))

	fmt.Println(anagram("aabbcc", "aabbc"))
}
