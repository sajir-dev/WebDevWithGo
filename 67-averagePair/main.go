package main

import "fmt"

func main() {
	a := []float64{1, 2, 6, 7, 8, 10}
	var avg float64 = 10
	fmt.Println(avgPair(a, avg))
}

func avgPair(arr []float64, avg float64) bool {
	// avgPair([1, 2, 6 ,7, 8, 10], 7) => true
	// avgPair([-1, 10, 100], 12) => false

	i := 0
	j := len(arr) - 1

	for i < j {
		sum := arr[j] + arr[i]
		if sum == 2*avg {
			return true
		} else if sum < 2*avg {
			i++
		} else {
			j--
		}
	}
	return false
}
