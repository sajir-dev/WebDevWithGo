package main

import "fmt"

func main() {
	arr := []int{10, 13, 1, 4, 2, -10}
	sum := 31
	fmt.Println(minSubArrayLen(arr, sum))
}

func minSubArrayLen(arr []int, num int) int {
	// take an array of positive integers
	// return min possible length of the subarray of which sum is greater than or equal to the given number
	// minSubArrayLen([1, 5, 8, 7, 2, 12], 13) => 2 ie; len([2, 12])
	// minSubArrayLen([12, 13, 14, 4, 2, 20], 31) => 3 ie; len([12, 13, 1] or [13, 14, 4])

	total := 0
	start := 0
	end := 0
	minLen := len(arr) + 100

	for start < len(arr) {
		if total < num && end < len(arr) {
			total += arr[end]
			end++
		} else if total >= num {
			minLen = end - start
			total -= arr[start]
			start++
		} else {
			break
		}
	}

	if minLen < len(arr) {
		return minLen
	}

	return 0

}

// JS solution
// function minSubArrayLen(nums, sum) {
// 	let total = 0;
// 	let start = 0;
// 	let end = 0;
// 	let minLen = Infinity;

// 	while (start < nums.length) {
// 		// if current window doesn't add up to the given sum then
// 		// move the window to right
// 		if(total < sum && end < nums.length){
// 		total += nums[end];
// 				end++;
// 		}
// 		// if current window adds up to at least the sum given then
// 		// we can shrink the window
// 		else if(total >= sum){
// 		minLen = Math.min(minLen, end-start);
// 				total -= nums[start];
// 				start++;
// 		}
// 		// current total less than required total but we reach the end, need this or else we'll be in an infinite loop
// 		else {
// 		break;
// 		}
// 	}

// 	return minLen === Infinity ? 0 : minLen;
// }
