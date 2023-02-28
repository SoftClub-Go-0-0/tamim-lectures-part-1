package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	arr[0], arr[1] = 1, 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}

}
