package main

import "fmt"

func main() {
	arr := []int{1, 45, 58, -10}
	fmt.Println(insSort(arr))
}

func insSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currentVal := arr[i]
		j := i - 1
		fmt.Println(arr)
		for j >= 0 && arr[j] > currentVal {
			arr[j+1] = arr[j]
			j--
			fmt.Println(arr)
		}
		arr[j+1] = currentVal
	}
	return arr
}
