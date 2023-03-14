package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		file.Write([]byte(strconv.Itoa(i*2) + " "))
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
