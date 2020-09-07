// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Shape represents...
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Rect represents...
// type Rect struct {
// 	width	float64
// 	height	float64
// }

// // Circle represents...
// type Circle struct {
// 	radius float64
// }

// // Area represents...
// func (r Rect) Area() float64 {
// 	return r.width * r.height
// }

// // Perimeter represents...
// func (r Rect) Perimeter() float64 {
// 	return 2 * (r.height + r.width)
// } 

// // Area represents...
// func (c Circle) Area() float64 {
// 	return math.Phi * c.radius * c.radius
// }

// // Perimeter represents...
// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Phi * c.radius
// }

// func main() {
// 	var s Shape = Rect{10, 3}

// 	fmt.Printf("type of s is %T\n", s)
// 	fmt.Printf("value of s is %v\n", s)
// 	fmt.Printf("value of s is %0.2v\n\n", s.Area())

// 	s = Circle{10}
// 	fmt.Printf("type of s is %T\n", s)
// 	fmt.Printf("value of s is %v\n", s)
// 	fmt.Printf("value of s is %0.2v\n\n", s.Area())
// }

// ---

// package main

// import (
// 	"fmt"
// )

// // MyString explains...
// type MyString string

// // Rect explains...
// type Rect struct {
// 	width	float64
// 	height	float64
// }

// func explain(i interface{}) {
// 	fmt.Printf("value given to explain function is of type '%T' with value of '%v'\n", i, i)
// }

// func main() {
// 	ms := MyString("Hello World!")
// 	r := Rect{5.5, 4.5}
// 	explain(ms)
// 	explain(r)
// }

// ---

// package main

// import (
// 	"fmt"
// )

// // Shape represents...
// type Shape interface{
// 	Area() float64
// }

// // Object represents...
// type Object interface{
// 	Volume() float64
// }

// // Skin represents...
// type Skin interface{
// 	Color() float64
// }

// // Cube represents...
// type Cube struct {
// 	side float64
// }

// // Area represents...
// func (c Cube) Area() float64 {
// 	return 6 * (c.side * c.side)
// }

// // Volume represents...
// func (c Cube) Volume() float64 {
// 	return c.side * c.side * c.side
// }

// func main() {
// 	var s Shape = Cube{3}
// 	value1, ok1 := s.(Object)
// 	fmt.Printf("dynamic value of Shape 's' with value %v implements interface object? %v\n", value1, ok1)
// 	value2, ok2 := s.(Skin)
// 	fmt.Printf("dynamic value of Shape 's' with value %v implements interface skin? %v\n", value2, ok2)
// }

// ---

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // Person interface
// type Person interface {
// 	getFullName() string
// }

// // Salaried interface
// type Salaried interface {
// 	getSalary() int
// }

// // Employee struct represents an employee in an origanization
// type Employee struct {
// 	firstName string
// 	lastName string
// 	salary int
// }

// // using this method, Employee implements Person interface
// func (e Employee) getFullName() string {
// 	return e.firstName + " " + e.lastName
// }

// // using this method, Employee implements Salaried interface
// func (e Employee) getSalary() int {
// 	return e.salary
// }


// func main() {
// 	var johnP Person = Employee{"John", "Adams", 2000}
	
// 	// show john's salary
// 	fmt.Printf("full name : %v \n", reflect.ValueOf(johnP).Interface())
	
// 	// convert john to Salaried type
// 	johnS := johnP.(Salaried)
	
// 	fmt.Printf("salary : %v \n", johnS.getSalary())

// }

// --- Type Switch

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func explain (i interface{}) {
// 	switch i.(type) {
// 	case string:
// 		fmt.Println("i stored string", strings.ToUpper(i.(string)))
// 	case int:
// 		fmt.Println("i stored int", i)
// 	default:
// 		fmt.Println("i stored something else", i)
// 	}
// }

// func main() {
// 	explain("Hello World")
// 	explain(52)
// 	explain(true)
// }

// --- Embedding Interface

// package main

// import "fmt"

// type Shape interface {
// 	Area() float64
// }

// type Object interface {
// 	Volume() float64
// }

// type Material interface {
// 	Shape
// 	Object
// }

// type Cube struct {
// 	side float64
// }

// func (c Cube) Area() float64 {
// 	return 6 * (c.side * c.side)
// }

// func (c Cube) Volume() float64 {
// 	return c.side * c.side * c.side
// }

// func main() {
// 	c := Cube{3}
// 	var m Material = c
// 	var s Shape = c
// 	var o Object = c
// 	fmt.Printf("dynamic type and value of interface m of static type Material is '%T' and '%v'\n", m, m)
// 	fmt.Printf("dynamic type and value of interface s of static type Shape is '%T' and '%v'\n",s, s)
// 	fmt.Printf("dynamic type and value of interface o of static type Object is '%T' and '%v'\n", o, o)
// }

// --- Pointer vs Value receiver

package main

import (
	"fmt"
)

type Shape interface {
	Area() 		float64
	Perimeter() float64
}

type Rect struct {
	width 	float64
	height 	float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	r := Rect{5.0, 4.0}
	var s Shape = &r
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Println("Area of rectangle is", area)
	fmt.Println("Perimeter of rectangle is", perimeter)
}