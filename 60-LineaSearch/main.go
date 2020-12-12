package main

import "fmt"

func main() {
	// fmt.Println(searchBinary([]int{23, 24, 35, 4}, 4))
	fmt.Println(searchBinary([]int{1, 3, 8, 10, 14}, 80))
}

// func searchLinear(s []int, key int) bool {
// 	for _, v := range s {
// 		// fmt.Println(v)
// 		if key == v {
// 			return true
// 		}
// 	}
// 	return false
// }

func searchBinary(s []int, key int) int {
	first := s[0]
	last := len(s) - 1
	mid := (len(s) - 1) / 2

	for {
		if key == s[mid] {
			return mid
		} else if first >= mid {
			return -1
		} else if key > s[mid] {
			first = mid
			mid = (first + last) / 2
		} else {
			last = mid
			mid = (first + last) / 2
		}
	}

}
