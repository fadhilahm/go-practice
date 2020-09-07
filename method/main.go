// package main

// import (
// 	"fmt"
// )

// // Employee stands for...
// type Employee struct {
// 	name string
// 	salary int
// }

// func (e *Employee) changeName( newName string) {
// 	e.name = newName
// }

// func main() {
// 	e := Employee{
// 		name: "Rose Geller",
// 		salary: 1200,
// 	}

// 	// e before the name change
// 	fmt.Println("e before name change =", e.name)

// 	// change name
// 	e.changeName("Monica Geller")

// 	// e after name change
// 	fmt.Println("e after name change =", e.name)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Rect is ...
// type Rect struct {
// 	width float64
// 	height float64
// }

// // Circle is ...
// type Circle struct {
// 	radius float64
// }

// // Area is ...
// func (r Rect) Area() float64 {
// 	return r.height * r.width
// }

// // Area is ...
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func main() {
// 	rect := Rect{5.0, 4.0}
// 	cir := Circle{5.0}
// 	fmt.Printf("Area of rectangle rect: %.2f\n", rect.Area())
// 	fmt.Printf("Area of circle cir: %.2f\n", cir.Area())
// }


// ---

// package main

// import (
// 	"fmt"
// )

// // Contact is ...
// type Contact struct {
// 	phone, address string
// }

// // Employee is ...
// type Employee struct {
// 	name 	string
// 	salary 	int
// }

// func (e *Employee) changeName(newName string) {
// 	e.name = newName
// }

// func (e Employee) showSalary() {
// 	e.salary = 1500
// 	fmt.Println("Salary of e =", e.salary)
// }

// func main() {
// 	e := Employee{
// 		name: "Ross Geller",
// 		salary: 1200,
// 	}

// 	// e before change
// 	fmt.Println("e before change =", e)

// 	// calling `changeName` pointer method on value
// 	e.changeName("Monica Geller")

// 	// calling `showSalary` value method on pointer
// 	(&e).showSalary()

// 	// e after change
// 	fmt.Println("e after change =", e)
// }

// ---

package main

import (
	"fmt"
	"strings"
)

// MyString is a derivation of `string` type
type MyString string

func (s MyString) toUpperCase () string {
	normalString := string(s)
	return strings.ToUpper(normalString)
}


func main() {
	str := MyString("Hello Happy World")
	fmt.Println("Things have changed:", str.toUpperCase())
}
