package main

import (
	"fmt"
)

func main() {
	var averageScore int
	fmt.Scan(&averageScore)
	switch {
	case averageScore < 80:
		fmt.Println("Sorry, you can not countinue the course")
	case averageScore > 100:
		fmt.Println("Hmm... I think you've made mistake. Enter the number from 0 to 100")
	default:
		fmt.Println("Congrates! You passed to the next stage.")
	}
}
