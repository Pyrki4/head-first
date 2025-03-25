package main

import (
	"fmt"
	"golang/chapter8/magazine"
)

func main() {
	employee := magazine.Employer{
		Name:   "Joy Carr",
		Salary: 60000,
		Address: magazine.Address{
			Street:     "123 Oak St",
			City:       "Omaha",
			State:      "NE",
			PostalCode: "68111"},
	}
	fmt.Println(employee.Name, employee.Salary, employee.City, employee.State, employee.Street, employee.PostalCode)
}
