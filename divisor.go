package main

import "fmt"

func divisor(n int) []int {
	sums := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sums = append(sums, i)
		}
	}
	fmt.Println(sums)
	return sums
}
