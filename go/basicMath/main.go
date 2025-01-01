package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	n := x
	reversed := 0
	for reversed < n {
		reversed = reversed*10 + n%10
		n /= 10
	}
	if reversed > n {
		return reversed/10 == n
	}
	return n == reversed
}

func countDigits() {
	n := 12
	x := n
	count := 0
	for x > 0 {
		y := x % 10
		if y != 0 {
			count = count + 1
		}
		y = x / 10
	}
}

func main() {
	fmt.Println(isPalindrome(11))
}
