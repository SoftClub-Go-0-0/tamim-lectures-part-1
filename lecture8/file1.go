package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println(false)
		return
	}
	fmt.Println(true)
	f.Close()
	os.Remove(f.Name())
}
