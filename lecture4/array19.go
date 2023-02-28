package main

import "fmt"

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	k := 9
	for arr[0] >= len(arr)-1 || len(arr)-1 >= arr[9] && k > 0 {
		k--
	}
	if k == 0 {
		fmt.Println(k)
	} else {
		fmt.Print(len(arr) - 1)
	}
}
