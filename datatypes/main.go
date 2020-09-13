// --- multiple variable declarations

package main

import (
	"fmt"
	"math"
)

// defining a type
// type float float64

// type alias
type float = float64

func main() {
	var (
		var1		= 50
		var2		= 25.22
		var3 string = "Telefonia"
		var4 bool
	)

	fmt.Println("var1 =", var1)
	fmt.Println("var2 =", var2)
	fmt.Println("var3 =", var3)
	fmt.Println("var4 =", var4)

	// multiple constant declarations
	const (
		consta = 1
		constb = 2
		constc
		constd
	)
	fmt.Println("consta =", consta)
	fmt.Println("constb =", constb)
	fmt.Println("constc =", constc)
	fmt.Println("constd =", constd)

	// iota expression
	const (
		iotaA = iota
		iotaB = iota
		iotaC = iota
		iotaD = iota + 1
		_ 	  = iota
		iotaE = iota
	)
	fmt.Println("iotaA =", iotaA)
	fmt.Println("iotaB =", iotaB)
	fmt.Println("iotaC =", iotaC)
	fmt.Println("iotaD =", iotaD)
	fmt.Println("iotaE =", iotaE)

	// more about short-hand notation
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("The minimum value is", c)

	// type conversion
	conv1 := 10
	conv2 := 10.5

	// illegal (mismatched types int and float64)
	// conv3 := conv1 + conv2

	// legal
	conv3 := conv1 + int(conv2)
	fmt.Println("conv3 =", conv3)

	// defining a type (*cont)
	var f float = 52.2
	var g float64 = 52.2
	fmt.Printf("f has value %v and type %T\n", f, f)
	fmt.Println("f == g", f == g)
}
