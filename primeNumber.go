package main

import "fmt"

func primeNumber(n int) bool {
	isPrimeNum := false
	sums := []int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j == n {
				sums = append(sums, i*j)
			}
		}
	}

	if len(sums) == 2 {
		isPrimeNum = true
	} else {
		isPrimeNum = false
	}
	fmt.Println(isPrimeNum)
	return isPrimeNum
}
