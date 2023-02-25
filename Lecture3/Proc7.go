package main

import "fmt"

func InvDigits(k *int) {
	ans := 0
	for i := *k; i != 0; i /= 10 {
		ans = ans*10 + i%10
	}
	k = &ans
}
func main() {
	var k int
	fmt.Scan(&k)
	InvDigits(&k)
	fmt.Println(k)
}
