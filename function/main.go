// # what is function

// package main

// import "fmt"

// func addMult(a, b int) (add, mul int) {
// 	add = a + b
// 	mul = a * b
// 	return
// }

// func main() {
// 	_, multRes := addMult(1, 5)
// 	fmt.Println(multRes)
// }

// # recursive

// package main

// import "fmt"

// // factorial ==> n! = n * (n - 1)! where n > 0
// func getFactorial(n int) int {
// 	if n > 1 {
// 		return n * getFactorial(n - 1)
// 	} 
// 		return 1
// }

// func main() {
// 	fmt.Println(getFactorial(5))
// }

// # defer

// package main

// import "fmt"

// func sayDone() {
// 	fmt.Println("I am done")
// }

// func main() {
// 	fmt.Println("main started")

// 	defer sayDone()

// 	fmt.Println("main finished")
// }

// package main

// import "fmt"

// func endTime(timestamp string){
// 	fmt.Println("Program ended at", timestamp)
// }

// func main() {
// 	time := "1 PM"

// 	defer endTime(time)

// 	time = "2 PM"

// 	fmt.Println("doing something")
// 	fmt.Println("main finished")
// 	fmt.Println("time is", time)
// }

// package main

// import "fmt"

// func greet (message string) {
// 	fmt.Println("greeting: ", message)
// }

// func main() {
// 	fmt.Println("Call one")

// 	defer greet("Greet one")

// 	fmt.Println("Call two")

// 	defer greet("Greet two")

// 	fmt.Println("Call three")

// 	defer greet("Greet three")
// }

// # function as type

// package main

// import "fmt"

// func add (a, b int) (c int) {
// 	c = a + b
// 	return
// }

// func substract (a, b int) (c int) {
// 	c = a - b
// 	return
// }

// func calc (a, b int, f func(int, int) int) (c int) {
// 	c = f(a,b)
// 	return
// }

// func main() {
// 	addResult := calc(5, 3, add)
// 	subResult := calc(5, 3, substract)
// 	fmt.Println("5 + 3 =", addResult)
// 	fmt.Println("5 - 3 =", subResult)
// }

// package main

// import "fmt"

// func add (a, b int) (c int) {
// 	c = a + b
// 	return
// }

// func substract (a, b int) (c int) {
// 	c = a - b
// 	return
// }

// // CalcFunc is..
// type CalcFunc func (int, int) int

// func calc(a, b int, f CalcFunc) (r int) {
// 	r = f(a, b)
// 	return
// }

// func main() {
// 	addResult := calc(5, 3, add)
// 	subResult := calc(5, 3, substract)
// 	fmt.Println("5 + 3 =", addResult)
// 	fmt.Println("5 - 3 =", subResult)
// }

// --- Function as a value
// package main

// import "fmt"

// var add = func (a int, b int) int {
// 	return a + b
// }

// func main() {
// 	fmt.Println("5 + 3 =", add(5, 3))
// }

// --- Immediately invoked function expression (IFE)
package main

import "fmt"

func main() {
	sum := func (a int, b int) int {
		return a + b
	}(3, 5)

	fmt.Println("5 + 3 =", sum)
}