package main

import "fmt"

func main() {
	var n int
	var a, d float64

	fmt.Scan(&n, &a, &d)

	aProgression := make([]float64, n)
	aProgression[0] = a

	for i := 1; i < n; i++ {
		aProgression[i] = aProgression[i-1] + d
	}

	fmt.Println(aProgression)

	// for i := 0; i < len(aProgression); i++ {
	// 	fmt.Printf("%.2f ", aProgression[i])
	// }

	// for _, i := range aProgression {
	// 	fmt.Printf("%.2f ", i)
	// }
	fmt.Println()
}
