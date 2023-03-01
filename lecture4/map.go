package main

import "fmt"

func main() {
	var taskNumber int
	tasks := make(map[string]bool)
	correctNumber, incorrectNumber := 0, 0
	var isSolved bool
	var taskName string
	fmt.Scan(&taskNumber)
	for i := 0; i < taskNumber; i++ {
		fmt.Scan(&taskName, &isSolved)
		tasks[taskName] = isSolved
		if isSolved {
			correctNumber++
		} else {
			incorrectNumber++
		}
	}
	fmt.Printf("Total correct tasks %d\n", correctNumber)
	fmt.Println("Correct tasks: ")
	for taskName, isSolved := range tasks {
		if isSolved {
			fmt.Printf("%s ", taskName)
		}
	}
	fmt.Println()

	fmt.Printf("Total incorrect tasks %d\n", incorrectNumber)
	fmt.Println("incorrect tasks: ")
	for taskName, isSolved := range tasks {
		if !isSolved {
			fmt.Printf("%s ", taskName)
		}
	}
}
