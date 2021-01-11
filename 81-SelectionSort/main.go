package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		index := i
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				index = j
			}
		}
		if i != index {
			arr[index] = arr[i]
			arr[i] = min
		}
	}
	return arr
}

func main() {
	arr := []int{-10, 8, -100, 200}
	fmt.Println(selectionSort(arr))
}
