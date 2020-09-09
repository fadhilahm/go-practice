// package main

// import (
// 	"fmt"
// )

// // Employee stands for a single employee object
// type Employee struct {
// 	firstName, lastName string
// 	Salary
// 	bool
// }

// // Data stands for justice
// type Data struct {
// 	string
// 	int
// 	bool
// }

// // Salary stands for ....
// type Salary struct {
// 	basic int
// 	insurance int
// 	allowance int
// }

// func main() {
// 	ross := &Employee{
// 		firstName: "Ross",
// 		lastName: "Bing",
// 		bool: true,
// 		Salary: Salary{ 1100, 50, 50 },
// 	}

// 	// fmt.Println("ross.firstName:", ross.firstName)
// 	// fmt.Println("ross.lastName:", ross.lastName)
// 	// fmt.Println("ross.salary:", ross.salary)
// 	// fmt.Println("ross.bool:", ross.bool)

// 	// sample1 := Data{ "Monday", 1200, true }
// 	// fmt.Println(sample1)
// 	// fmt.Println(ross)

// 	fmt.Println("Ross basic salary:", ross.basic)
// }

// --- nested interface
// package main

// import (
// 	"fmt"
// )

// type Salaried interface {
// 	getSalary() int
// }

// type Salary struct {
// 	basic		int
// 	insurance	int
// 	allowance 	int
// }

// func (s Salary) getSalary() int {
// 	return s.basic + s.insurance + s.allowance
// }

// type Employee struct {
// 	firstName, lastName string
// 	salary Salaried
// }

// func main() {
// 	ross := Employee{
// 		firstName: 	"Ross",
// 		lastName: 	"Geller",
// 		salary:		Salary{1100, 50, 50},
// 	}

// 	fmt.Println("Ross's salary is", ross.salary.getSalary())
// }

// --- promoted interface
// package main

// import (
// 	"fmt"
// )

// type Salaried interface {
// 	getSalary() int
// }

// type Salary struct {
// 	basic		int
// 	insurance	int
// 	allowance 	int
// }

// func (s Salary) getSalary() int {
// 	return s.basic + s.insurance + s.allowance
// }

// type Employee struct {
// 	firstName, lastName string
// 	Salaried
// }

// func main() {
// 	ross := Employee{
// 		firstName: 	"Ross",
// 		lastName: 	"Geller",
// 		Salaried:	Salary{1100, 50, 50},
// 	}

// 	fmt.Println("Ross's salary is", ross.getSalary())
// }

// --- Exporting struct
// package main

// import (
// 	"fmt"
// 	"org"
// )

// type Employee org.Employee

// func main() {
// 	ross := org.Employee{
// 		FirstName: "Ross",
// 		LastName: "Geller",
// 	}

// 	fmt.Println(ross)
// }

// --- Function fields

// package main

// import "fmt"

// type FullNameType func (string, string) string

// type Employee struct {
// 	Firstname, Lastname 	string
// 	Fullname				FullNameType
// }

// func main() {
// 	e := Employee{
// 		Firstname:	"Ross",
// 		Lastname:	"Geller",
// 		Fullname:	func(firstName string, lastName string) string {
// 			return firstName + " " + lastName
// 		},
// 	}

// 	fmt.Println(e.Fullname(e.Firstname, e.Lastname))
// }

// --- Struct comparison

package main

import "fmt"

// Employee is....
type Employee struct {
	FirstName		string	`json:"firstName"`
	LastName		string	`json:"lastName"`
	Salary			int		`json:"salary"`
}

func main() {
	ross := Employee{
		FirstName: "Ross",
		LastName: "Geller",
		Salary: 1200,
	}

	rossCopy := Employee{
		FirstName: "Ross",
		LastName: "Geller",
		Salary: 1200,
	}

	fmt.Println(ross == rossCopy)
}

