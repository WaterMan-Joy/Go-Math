package main

import "fmt"

func multiple(n int, count int) []int {
	sums := []int{}
	for i := 1; i <= count; i++ {
		if i <= count {
			sums = append(sums, n*i)
		}
	}
	fmt.Println(sums)
	return sums
}
