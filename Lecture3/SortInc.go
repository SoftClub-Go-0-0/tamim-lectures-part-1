package main

import "fmt"

func SortInc(a, b, c *float64) {
	var max, min float64
	if *a > *b && *a > *c {
		max = *a
	} else if *b > *a && *b > *c {
		max = *b
	} else if *c > *b && *c > *a {
		max = *c
	}
	if *a < *b && *a < *c {
		min = *a
	} else if *b < *a && *b < *c {
		min = *b
	} else if *c < *b && *c < *a {
		min = *c
	}

	if min == *a && max == *c {
		fmt.Println(min, *b, max)
	} else if min == *a && max == *b {
		fmt.Println(min, *c, max)
	} else if min == *b && max == *a {
		fmt.Println(min, *c, max)
	} else if min == *b && max == *c {
		fmt.Println(min, a, max)
	} else if min == *c && max == *a {
		fmt.Println(min, *b, max)
	} else if min == *c && max == *b {
		fmt.Println(min, a, max)
	}
}
func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	SortInc(&a, &b, &c)
}
