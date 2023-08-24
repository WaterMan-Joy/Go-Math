package main

import "fmt"

func primeNumber(n int) bool {
	isPrimeNumber := false
	sums := []int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j == n {
				sums = append(sums, i*j)
			}
		}
	}

	if len(sums) == 2 {
		isPrimeNumber = true

	} else {
		isPrimeNumber = false
	}
	fmt.Println(isPrimeNumber)
	return isPrimeNumber
}
