package main

import "fmt"

func main() {
	arr := []int{12, 4, 4, 8, 90, -120, 10}
	fmt.Println(mergeSort(arr))
	// arr2 := []int{-100, 30}
	// fmt.Println(mergeArr(arr1, arr2))
}

// var i int

func mergeSort(arr []int) []int {
	// i++
	// fmt.Println(i)
	if len(arr) == 1 {
		return arr
	}
	left := mergeSort(arr[:(len(arr) / 2)])
	right := mergeSort(arr[(len(arr) / 2):])
	fmt.Println(left, right)
	return mergeArr(left, right)
}

func mergeArr(arr1 []int, arr2 []int) []int {
	// 1, 6, 9, 12 i
	//  21, 81, 181, 211 j
	// compare arr1[i] and arr2[j]
	newArr := make([]int, len(arr1)+len(arr2))
	i := 0
	j := 0
	for i+j < len(arr1)+len(arr2) {
		if i == len(arr1) && j != len(arr2) {
			newArr[i+j] = arr2[j]
			j++
			continue
		}
		if i != len(arr1) && j == len(arr2) {
			newArr[i+j] = arr1[i]
			i++
			continue
		}
		if arr1[i] <= arr2[j] {
			newArr[i+j] = arr1[i]
			i++
		} else {
			newArr[i+j] = arr2[j]
			j++
		}
	}
	return newArr
}
