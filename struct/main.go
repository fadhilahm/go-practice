package main

import (
	"fmt"
)

// Employee stands for a single employee object
type Employee struct {
	firstName, lastName string
	Salary
	bool
}

// Data stands for justice
type Data struct {
	string
	int
	bool
}

// Salary stands for ....
type Salary struct {
	basic int
	insurance int
	allowance int
}

func main() {
	ross := &Employee{
		firstName: "Ross",
		lastName: "Bing",
		bool: true,
		Salary: Salary{ 1100, 50, 50 },
	}

	// fmt.Println("ross.firstName:", ross.firstName)
	// fmt.Println("ross.lastName:", ross.lastName)
	// fmt.Println("ross.salary:", ross.salary)
	// fmt.Println("ross.bool:", ross.bool)

	// sample1 := Data{ "Monday", 1200, true }
	// fmt.Println(sample1)
	// fmt.Println(ross)

	fmt.Println("Ross basic salary:", ross.basic)
}