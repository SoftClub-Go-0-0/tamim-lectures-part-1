package main

import "fmt"

type Person struct {
	name, surname string
	age           int
	heght, weight float64
	gender        string
}

func main() {
	var person Person
	person.name = "Tamim"
	person.surname = "Orif"
	person.age = 19
	person.gender = "Male"
	person.heght, person.weight = 179.1, 63.5
	fmt.Println(person)
}
