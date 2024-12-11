package main

import "fmt"

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / pow(x, -n)
	}
	half := pow(x, n/2)
	if (n & 1) == 0 {
		return half * half
	}
	return half * half * x
}

func main() {
	fmt.Println(pow(2, 2))
}
