package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	arr := make([]int, n)

	arr[0], arr[1] = a, b

	for i := 2; i < n; i++ {
		arr[i] = 0
		for j := 0; j < i; j++ {
			arr[i] += arr[j]
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d : %d\n", i, arr[i])
	}
}
