package main

import "fmt"

func AddRightDigit(k *int) {
	ans := 0
	for i := *k; i != 0; i /= 10 {
		ans *= 10
		ans += i % 10
	}
	*k = ans
}
func main() {
	var k int
	fmt.Scan(&k)
	AddRightDigit(&k)
	fmt.Println()
}
