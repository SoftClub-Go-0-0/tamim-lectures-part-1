package main

import (
	"fmt"
	"math"
)

// AMean = (X+Y)/2
// GMean = pow((X·Y),1/2)
// (A, B), (A, C), (A, D)
func Mean(x, y float64, AMean, GMean *float64) {
	*AMean = (x + y) / 2
	*GMean = math.Pow((x * y), 1/2)
}
func main() {
	var a, b, c, d, AMean, GMean float64
	fmt.Scan(&a, &b, &c, &d)
	Mean(a, b, &AMean, &GMean)
	fmt.Println(AMean, " ", GMean)

	Mean(a, c, &AMean, &GMean)
	fmt.Println(AMean, " ", GMean)

	Mean(a, d, &AMean, &GMean)
	fmt.Println(AMean, " ", GMean)
}
