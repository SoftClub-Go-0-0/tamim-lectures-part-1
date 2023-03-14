package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var filename string

	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		stringNumber := scanner.Text()
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, number)
	}

	fmt.Println(numbers)
}
