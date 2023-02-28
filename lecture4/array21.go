package main

import "fmt"

func main() {
	var n, l, k int
	sum, ave := 0, 0
	fmt.Scanf("%d %d %d", &n, &k, &l)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := k; i < l; i++ {
		sum += arr[i]
		ave++
	}

	fmt.Println(float64(sum / ave))
}
