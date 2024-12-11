package main

import (
	"fmt"
)

func reverseArray(array []int, start int, end int) {
	if start >= end {
		return
	}
	t := array[start]
	array[start] = array[end]
	array[end] = t
	reverseArray(array, start+1, end-1)
}

func palindrome(array []string, i int, n int) bool {
	if i >= n/2 {
		return true
	}
	if array[i] != array[n-i-1] {
		return false
	}
	return palindrome(array, i+1, n)
}

func fib(x int) int {
	a := 0
	b := 1
	if x == a {
		return a
	}
	if x == b {
		return b
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println(fib(4))
}
