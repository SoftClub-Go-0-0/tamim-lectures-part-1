package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, _ := os.Create(filename)

	var a, d float64
	fmt.Scan(&a, &d)

	for i := 0; i < 10; i++ {
		currentElement := a + d*float64(i)

		currentElementString := strconv.FormatFloat(currentElement, 'f', 2, 64)

		_, err := file.WriteString(currentElementString + " ")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file.Close()
}
