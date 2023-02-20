package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b float64
	fmt.Scan(&n, &a, &b)
	d := math.Abs(a - b)
	each := d / n
	for i := 0; i < int(n); i++ {
		fmt.Println(1 - math.Sin(a+(each*float64(i))))
	}
	// Still got a error
}
