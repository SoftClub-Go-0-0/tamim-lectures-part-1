package main

import (
	"fmt"
	"math"
)

func TrianglePS(a float64, p, s *float64) {
	*p = 3 * a
	*s = (a * a) * math.Sqrt(3) / 4
}
func main() {
	var a, s, p float64
	fmt.Scan(&a)
	TrianglePS(a, &p, &s)
	fmt.Printf("%.2f\t%.2f", p, s)
}
