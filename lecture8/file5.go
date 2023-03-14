package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(-1)
		return
	}

	number := make([]byte, 300)
	file.Read(number)
	fmt.Println(string(number))

	u := User{}
	fmt.Scan(&u.Age)
}

type User struct {
	Age int
}
