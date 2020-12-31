package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 8, 1, 19}
	c := 2
	fmt.Println(maxSubArraySum(a, c))
}

func maxSubArraySum(arr []int, count int) int {
	// [1, 2, 3, 8, 1, 19], 2 => 20
	// [1,2, 8, 5, 11, 12, 1], 3 => 28
	sum := 0
	temp := 0
	for j := 0; j < count; j++ {
		temp = temp + arr[j] // 3
	}

	fmt.Println(temp)

	for j := count; j < len(arr); j++ {
		// fmt.Println(arr[count], arr[j-count])
		temp = temp + arr[j] - arr[j-count]
		fmt.Println(temp, arr[j], arr[j-count])
		if temp > sum {
			sum = temp
		}
	}

	return sum
}
