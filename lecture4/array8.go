package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])

	}
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			fmt.Print(arr[i], " ")
			k++
		}
	}
	fmt.Printf("\n%d\n", k)
}
