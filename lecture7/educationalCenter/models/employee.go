package models

import (
	"math/rand"
	"time"
	"educationalCenter/models"
)

type Employee struct {
	ID            int
	CompanyID     int
	Name          string
	Surname       string
	Email         string
	DateOfJoining time.Time
	Position      string
	Salary        int
}

func (e *Employee) TotalMoneyEarned() int {
	totalMonths := time.Now().Sub(e.DateOfJoining).Hours() / 24 / 30
	return int(totalMonths) * e.Salary
}

func NewEmployee(companyID int) Employee {
	var employee Employee

	employee.ID = rand.Intn(2000) + 1
	employee.CompanyID = companyID
	employee.Name = RandStringBytes(5)
	employee.Surname = RandStringBytes(15)
	employee.Email = RandStringBytes(10)

	years := rand.Intn(5)
	months := rand.Intn(12)

	employee.DateOfJoining = time.Now().AddDate(-years, -months, 0)
	employee.Position = RandStringBytes(10)
	employee.Salary = rand.Intn(6000) + 1 + 1000

	return employee
}
