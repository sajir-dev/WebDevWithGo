package main

import (
	"errors"
	"fmt"
)

func main() {
	a := []int{-5, -3, 0, 1, 2, 3, 4}
	b := []int{-5, 0, 1, 2, 3, 4}
	fmt.Println(sumZero(a))
	fmt.Println(sumZero(b))

	c := []int{1, 1, 1, 3, 5, 5, 8, 8, 8}
	d := []int{}
	fmt.Println(countUniqueValues1(c))
	fmt.Println(countUniqueValues2(c))
	fmt.Println(countUniqueValues1(d))
	fmt.Println(countUniqueValues2(d))
}

func sumZero(arr []int) ([]int, error) {
	// sumZero([-5, -3, 0, 1, 2, 3, 4]) should return [-3, 3]
	// sumZer0([-5, 0, 1, 2, 3, 4]) should return an error

	i := 0

	for j := len(arr) - 1; j > i; {
		k := arr[i] + arr[j]
		// fmt.Println(k, i, j)
		if k == 0 {
			return []int{arr[i], arr[j]}, nil
		} else if k > 0 {
			j--
		} else {
			i++
		}
	}

	return nil, errors.New("no pair gives sum zero")
}

func countUniqueValues1(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	return len(m)
}

func countUniqueValues2(arr []int) int {
	// [1, 1, 1, 3]
	if len(arr) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}
	return i + 1
}
