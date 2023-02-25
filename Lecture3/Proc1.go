package main

import (
	"fmt"
	"math"
)

func PowerA3(a float64, b *float64) {
	*b = math.Pow(a, 3)
}
func main() {
	var a, b float64
	fmt.Scan(&a)
	PowerA3(a, &b)
	fmt.Println(b)
}
