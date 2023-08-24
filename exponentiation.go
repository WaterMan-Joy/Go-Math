package main

import "fmt"

func exponentiation(n int, count int) int {
	sum := 1

	for i := 1; i <= count; i++ {
		sum = sum * n

	}
	fmt.Println(sum)
	return sum
}
