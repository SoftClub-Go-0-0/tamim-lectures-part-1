package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	pows := 1
	powers := make([]int, n)
	for i := 0; i < n; i++ {
		powers[i] = int(math.Pow(float64(pows), 2))
		pows++
	}
	fmt.Println(powers)
}
