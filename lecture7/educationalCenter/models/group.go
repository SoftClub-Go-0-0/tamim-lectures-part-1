package models

type Group struct {
	ID       uint
	Title    string
	Tutor    Person
	Students []Student
}
