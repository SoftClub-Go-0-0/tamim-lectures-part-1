package models
import (
	"fmt"
	"math/rand"
	"educationalCenter/models"
)

type Company struct {
	ID                 int
	Name               string
	Location           string
	YearOfFoundation   int
	CurrentYearIncome  int
	PreviousYearIncome int
	Employees          []Employee
}

// TotalMoneySpentOnSalaries возвращает общее количество денег потраченных на оплату труда сотрудников
func (c *Company) TotalMoneySpentOnSalaries() (totalMoney int) {

	return
}

// YearsOfExistence возвращает разницу между текущим годом и годом основания компании
func (c *Company) YearsOfExistence() int {
	return 2023 - c.YearOfFoundation
}

// IncomeDifference возвращает разницу между доходом текущего и предыдущего года
func (c *Company) IncomeDifference() int {
	return c.CurrentYearIncome - c.PreviousYearIncome
}

// HasEmployee принимает имя сотрудника и проверяет есть ли такой сотрудник в массиве сотрудников компании
func (c *Company) HasEmployee(name string) bool {
	for _, employee := range c.Employees {
		if employee.Name == name {
			return true
		}
	}
	return false
}

// implement employee

func (c *Company) PrintInfo() {
	fmt.Printf("ID:\t\t\t%d\n", c.ID)
	fmt.Printf("Name:\t\t\t%s\n", c.Name)
	fmt.Printf("Location:\t\t%s\n", c.Location)
	fmt.Printf("Year of foundation:\t%d\n", c.YearOfFoundation)
	fmt.Printf("Current year income:\t%d\n", c.CurrentYearIncome)
	fmt.Printf("Previous year income:\t%d\n", c.PreviousYearIncome)
	fmt.Println("Employees")
	for _, e := range c.Employees {
		fmt.Println(e)
	}
}

func NewCompany() Company {
	var company Company

	company.ID = rand.Intn(60) + 1
	company.Name = RandStringBytes(5)
	company.Location = RandStringBytes(10)
	company.YearOfFoundation = rand.Intn(1000) + 1000
	company.CurrentYearIncome = rand.Intn(1000000) + 1
	company.PreviousYearIncome = rand.Intn(1000000) + 1

	employeesNumber := rand.Intn(50) + 1

	company.Employees = make([]Employee, employeesNumber)
	for i := 0; i < employeesNumber; i++ {
		company.Employees[i] = NewEmployee(company.ID)
	}
	return company
}
