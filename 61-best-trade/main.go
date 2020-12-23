package main

import "fmt"

func main() {
	arr := []int{25, 35, 200, 56, 2, 45, 11}
	fmt.Println(bestTrade(arr))
}

func bestTrade(arr []int) (int, int, int) {
	maxProfit, buy, sell := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] > maxProfit {
				maxProfit = arr[j] - arr[i]
				buy = i
				sell = j
			}
		}
	}
	return maxProfit, buy, sell
}

// a = [25, 35, 42, 56, 2, 45, 11]
// find the biggest difference
