package models

type Student struct {
	Person         Person
	GroupID        uint
	Grades         map[string]int
	TopicsToWorkOn []string
}
