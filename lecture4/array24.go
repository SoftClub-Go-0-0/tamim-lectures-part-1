package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	x := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if x != arr[i]-arr[i-1] {
			x = 0
		}
	}
	fmt.Println(x)
}
