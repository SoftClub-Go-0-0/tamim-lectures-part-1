package main

import (
	"fmt"
	"math"
)

func PowerA234(a float64, b, c, d *float64) {
	*b = math.Pow(a, 2)
	*c = math.Pow(a, 3)
	*d = math.Pow(a, 4)
}
func main() {
	var a, b, c, d float64
	fmt.Scan(&a)
	PowerA234(a, &b, &c, &d)
	fmt.Print(b, " ")
	fmt.Print(c, " ")
	fmt.Println(d)
}
