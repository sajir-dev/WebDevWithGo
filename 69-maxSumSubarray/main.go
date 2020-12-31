package main

import "fmt"

func main() {
	var a = []float64{12, 1, 5, 8, 12, 19, 4}
	count := 9
	fmt.Println(maxSumSubarray(a, count))
}

func maxSumSubarray(arr []float64, count int) float64 {
	if count >= len(arr) {
		return 0.0
	}
	sum := 0.00
	for j := 0; j < count; j++ {
		sum += arr[j]
	}

	temp := sum

	for j := count; j < len(arr); j++ {
		temp = temp - arr[j-count] + arr[j]
		fmt.Println(sum, temp, arr[j-count], arr[j])
		if temp > sum {
			sum = temp
		}
	}

	return sum
}
