package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	x := arr[0] % 2
	for i := 1; i < len(arr) && x != arr[i]%2; i++ {
		x = arr[i] % 2
	}
	if x == 0 {
		fmt.Println("0")
	} else {
		fmt.Print(arr[0])
	}
}
