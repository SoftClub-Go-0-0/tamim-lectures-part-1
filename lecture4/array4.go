package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var a, d float64
	fmt.Scan(&n, &a, &d)

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		arr[i] = a * math.Pow(d, float64(i))
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
