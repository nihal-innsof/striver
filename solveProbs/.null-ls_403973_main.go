package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	for k, v := range slices.Concat(nums[k:], nums[:k]) {
		nums[k] = v
	}
}

// First approach
/* func moveZeroes(nums []int) {
	p1 := 0
	temp := 0
	for i := 1; i < len(nums); i++ {
		if nums[p1] != 0 {
			p1 += 1
			continue
		}
		if nums[p1] == 0 && nums[p1] != nums[i] {
			temp = nums[p1]
			nums[p1] = nums[i]
			nums[i] = temp
			p1 += 1
		}
	}
} */

// Optimal approach
func moveZeroes(nums []int) {
	l := 0
	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l += 1
		}
	}

	for ; l < len(nums); l++ {
		nums[l] = 0
	}
}

func linearSearch() {
}

func main() {
	// data := []int{1, 2, 3, 4, 5, 6, 7}
	// data := []int{-1, -100, 3, 99}
	// rotate(data, 3)
	// data := []int{0, 1, 0, 3, 12}
	// data := []int{1, 0, 1}
	data := []int{45192, 0, -659, -52359, -99225, -75991, 0, -15155, 27382, 59818, 0, -30645, -17025, 81209, 887, 64648}
	moveZeroes(data)
	fmt.Println(data)
}
