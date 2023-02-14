package main

import "fmt"

func main() {
	// var name string
	// fmt.Print("Enter your ma,e, please: ")
	// fmt.Scan(&name)
	// println("Welcome to golang course", name)

	var age, phone_number, salary int
	var name,surname string
	fmt.Print("What is your name ? ")
	fmt.Scanf("%s", &name)

	fmt.Print("What is your surname ? ")
	fmt.Scanf("%s", &surname)

	fmt.Print("How old are you ? ")
	fmt.Scanf("%d", &age)

	fmt.Print("how much is your salary ? ")
	fmt.Scanf("%d", &salary)

	fmt.Print("What is your phone number ? ")
	fmt.Scanf("%d", &phone_number)

	fmt.Println("Welcome Mr/Miss",name,surname,"Your are",age,"and your phone number is",phone_number)
}
