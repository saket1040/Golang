package main

import "fmt"

type Employee interface {
	GetDetails() string
	CalculateSalary() float64
}

type FullTimeEmployee struct {
	Name       string
	BaseSalary float64
}

func (f *FullTimeEmployee) GetDetails() string {
	return f.Name
}

func (f *FullTimeEmployee) CalculateSalary() float64 {
	return f.BaseSalary
}

type ContractEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked int32
}

func (c *ContractEmployee) GetDetails() string {
	return c.Name + " worked at rate " + fmt.Sprintf("%.2f", c.HourlyRate) + " for " + fmt.Sprintf("%d", c.HoursWorked) + " hours"
}

func (c *ContractEmployee) CalculateSalary() float64 {
	return c.HourlyRate * float64(c.HoursWorked)
}

func PrintSalaryDetails(employees []Employee) {
	for _, emp := range employees {
		fmt.Println("Details:", emp.GetDetails())
		fmt.Println("Salary:", emp.CalculateSalary())
	}
}

func main() {
	fullTimeEmp := &FullTimeEmployee{Name: "John Doe", BaseSalary: 50000}
	contractEmp := &ContractEmployee{Name: "Jane Smith", HourlyRate: 50, HoursWorked: 160}

	employees := []Employee{fullTimeEmp, contractEmp}
	PrintSalaryDetails(employees)
}
