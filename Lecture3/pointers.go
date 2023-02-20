package main

import "fmt"

func main() {
	a := 10
	ptr := &a
	fmt.Printf("Value is %d and address is %p\n", a, ptr)
}
