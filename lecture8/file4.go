package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	var cnt int

	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err == nil {
		cnt++
	}
	file.Close()

	fmt.Scan(&filename)

	file, err = os.Open(filename)
	if err == nil {
		cnt++
	}
	file.Close()

	fmt.Scan(&filename)

	file, err = os.Open(filename)
	if err == nil {
		cnt++
	}
	file.Close()

	fmt.Scan(&filename)

	file, err = os.Open(filename)
	if err == nil {
		cnt++
	}
	file.Close()

	fmt.Println(cnt)
}
