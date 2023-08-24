package main

import (
	"fmt"
	"strconv"
)

func binary(text string) {
	for _, t := range text {
		binary := strconv.FormatInt(int64(t), 2)
		fmt.Printf("%s ", binary)
	}
}
