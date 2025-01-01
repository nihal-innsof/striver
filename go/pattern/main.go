package main

import "fmt"

func pattern7() {
	count := 4
	for i := 1; i <= count; i++ {
		for j := 1; j <= 2*count-1; j++ {
			if (j >= (count - i + 1)) && (j <= (count + i - 1)) {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern8() {
	count := 6
	for i := 1; i <= count; i++ {
		for j := 1; j <= 2*count-1; j++ {
			if (j >= (count - (count - i))) && (j <= (count + (count - i))) {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern9() {
	count := 4
	for i := 1; i <= count; i++ {
		for j := 1; j <= 2*count-1; j++ {
			if (j >= (count - i + 1)) && (j <= (count + i - 1)) {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for i := 1; i <= count; i++ {
		for j := 1; j <= 2*count-1; j++ {
			if (j >= (count - (count - i))) && (j <= (count + (count - i))) {
				fmt.Print("*")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern10() {
	n := 10
	for i := 1; i <= (2*n)-1; i++ {
		for j := 1; j <= i; j++ {
			if i <= n {
				fmt.Print("*")
			} else if j <= n-(i%n) {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}

func pattern11() {
	n := 5
	var pattern int
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			pattern = 1
		} else {
			pattern = 0
		}
		for j := 1; j <= i; j++ {
			fmt.Print(pattern)
			pattern = 1 - pattern
		}
		fmt.Println()
	}
}

func pattern12() {
	n := 4
	for i := 1; i <= n; i++ {
		spaceCount := ((2 * n) - (2 * i))
		numberCount := n - (spaceCount / 2)
		for j := 1; j <= numberCount; j++ {
			fmt.Print(j)
		}
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := numberCount; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func pattern13() {
	n := 5
	value := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(value)
			value = value + 1
		}
		fmt.Println()
	}
}

func pattern14() {
	n := 5
	startingCode := 65
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingCode + j - 1)))
		}
		fmt.Println()
	}
}

func pattern15() {
	n := 5
	startingCode := 65
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingCode + j - 1)))
		}
		fmt.Println()
	}
}

func pattern16() {
	n := 5
	startingCode := 65
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingCode + i - 1)))
		}
		fmt.Println()
	}
}

func pattern17() {
	n := 26
	startingCharacter := 'A'
	startingCode := int(startingCharacter) - n - 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2*n-1; j++ {
			if j >= (n-i+1) && j <= (n+i-1) {
				if j <= n {
					fmt.Print(string(rune(startingCode + i + j)))
				} else {
					fmt.Print(string(rune(startingCode + i + j - (j%n)*2)))
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func pattern18() {
	n := 5
	startingCharacter := 'A'
	startingCode := int(startingCharacter)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingCode + n - i + j - 1)))
		}
		fmt.Println()
	}
}

func pattern19() {
	n := 30
	for i := n; i >= 1; i-- {
		spaceCount := ((2 * n) - (2 * i))
		numberCount := n - (spaceCount / 2)
		for j := 1; j <= numberCount; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := numberCount; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= n; i++ {
		spaceCount := ((2 * n) - (2 * i))
		numberCount := n - (spaceCount / 2)
		for j := 1; j <= numberCount; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := numberCount; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern20() {
	n := 6
	for i := 1; i <= n; i++ {
		spaceCount := ((2 * n) - (2 * i))
		numberCount := n - (spaceCount / 2)
		for j := 1; j <= numberCount; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := numberCount; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		spaceCount := ((2 * n) - (2 * i))
		numberCount := n - (spaceCount / 2)
		for j := 1; j <= numberCount; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := numberCount; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern21() {
	n := 50
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (j == 1) || (i == 1) || (j == n) || (i == n) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func pattern22() {
	n := 3
	for i := 1; i <= 2*n-1; i++ {
		for j := 1; j <= 2*n-1; j++ {
			/* if (i == 1) || (j == 1) || (i == 2*n-1) || (j == 2*n-1) {
				fmt.Print(n)
			} else if (i == 2) || (j == 2) || (i == 2*n-2) || (j == 2*n-2) {
				fmt.Print(n - 1)
			} else if (i == 3) || (j == 3) || (i == 2*n-3) || (j == 2*n-3) {
				fmt.Print(n - 2)
			} else {
				fmt.Print(" ")
			} */
		}
		fmt.Println()
	}
}

func main() {
	pattern22()
}
