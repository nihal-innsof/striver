package main

import (
	"fmt"
	"slices"
)

// Bruteforce approach. Time complexity: O(n + n/2)
// func rearrangeArray(nums []int) []int {
// 	a := nums
// 	var p, n, result []int
// 	for i := 0; i < len(a); i++ {
// 		if a[i] > 0 {
// 			p = append(p, a[i])
// 		} else {
// 			n = append(n, a[i])
// 		}
// 	}
//
// 	for i := 0; i < len(a)/2; i++ {
// 		result = append(result, p[i])
// 		result = append(result, n[i])
// 	}
// 	return result
// }

// func rearrangeArray(nums []int) []int {
// 	a := nums
// 	var result []int
// 	var queue []int
// 	for i := 0; i < len(a); i++ {
// 		x := len(result)
// 		if isEven(x) && a[i] < 0 {
// 			// temp = a[i]
// 			queue = append(queue, a[i])
// 			continue
// 		} else if isOdd(x) && a[i] > 0 {
// 			// temp = a[i]
// 			queue = append(queue, a[i])
// 			continue
// 		}
// 		result = append(result, a[i])
// 		if len(queue) != 0 {
// 			result = append(result, queue[0])
// 			queue = queue[1:]
// 		}
// 	}
// 	return result
// }

// func isEven(x int) bool {
// 	return x%2 == 0
// }
//
// func isOdd(x int) bool {
// 	return x%2 != 0
// }

func rearrangeArray(nums []int) []int {
	a := nums
	var result []int
	var pi, ni int
	pi = 0
	ni = 1
	result = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			result[pi] = a[i]
			pi += 2
		} else {
			result[ni] = a[i]
			ni += 2
		}
		if ni > len(a) || pi > len(a) {
			break
		}
	}
	return result
}

// Bruteforce approach. TC O(n**2)
// func leaders(nums []int) []int {
// 	a := nums
// 	var result []int
// 	var skipFlag bool
// 	for i := 0; i < len(a); i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if a[i] < a[j] {
// 				skipFlag = true
// 				break
// 			}
// 		}
// 		if !skipFlag {
// 			result = append(result, a[i])
// 		}
// 		skipFlag = false
// 	}
// 	return result
// }

// Optimal approach: O(n)
func leaders(nums []int) []int {
	a := nums
	var result []int
	result = append(result, a[len(a)-1])
	var largest int
	largest = result[0]
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > largest {
			largest = a[i]
			result = append(result, largest)
		}
	}
	slices.Reverse(result)
	return result
}

func main() {
	var x []int
	// x = []int{3, 1, -2, -5, 2, -4}
	// x = []int{1, 2, -4, -5}
	// x = []int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}
	// fmt.Println(rearrangeArray(x))
	// x = []int{4, 7, 1, 0}
	x = []int{10, 22, 12, 3, 0, 6}
	fmt.Println(leaders(x))
}
