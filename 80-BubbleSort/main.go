package main

import "fmt"

func main() {
	arr := []int{2, 34, 100, 54, 10}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		noSwaps := true
		for j := 0; j < len(arr)-i-1; j++ {
			fmt.Println(j, arr)
			if arr[j] < arr[j+1] {
				// fmt.Println(j, arr)
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				// arr[j+1], arr[j] = arr[j+1], arr[j]
				noSwaps = false
				fmt.Println(j, arr)
			}
		}
		if noSwaps {
			break
		}
	}
	return arr
}
