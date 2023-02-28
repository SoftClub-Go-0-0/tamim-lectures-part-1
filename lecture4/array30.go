package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	x := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			x++
		}
	}
	fmt.Println(x)
}
