package main

import (
	"fmt"
	"math"
)

func RectPS(x1, x2, y1, y2 float64, p, s *float64) {
	a := math.Abs(x1 - x2)
	b := math.Abs(y1 - y2)
	*p = 2 * (a + b)
	*s = a * b
}
func main() {
	var x1, x2, y1, y2, s, p float64
	fmt.Scan(&x1, &x2, &y1, &y2)
	RectPS(x1, x2, y1, y2, &s, &p)
	fmt.Printf("%.2f\t%.2f", p, s)
}
