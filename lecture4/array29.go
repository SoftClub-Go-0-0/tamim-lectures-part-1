package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	x := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]%2 == 1 && arr[i] > x {
			x = arr[i]
		}
	}
	fmt.Println(x)
}
