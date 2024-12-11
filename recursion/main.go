package main

import "fmt"

const MOD = int(1e9 + 7)

func binexp(a, b, MOD int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a)
		}
		a = (a * a)
		b >>= 1
	}

	return res % MOD
}

func countGoodNumbers(n int64) int {
	odds := int(n / 2)
	evens := int((n + 1) / 2)
	return (binexp(5, evens, MOD) * binexp(4, odds, MOD)) % MOD
}

func main() {
	fmt.Println(countGoodNumbers(50))
}
