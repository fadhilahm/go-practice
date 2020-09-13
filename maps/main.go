// --- what is a Map?
// package main

// import "fmt"

// func main() {
// 	var m map[string] int

// 	fmt.Println(m)
// 	fmt.Println("m == nil", m == nil)
// }

// --- creating a simple map
// package main

// import "fmt"

// func main() {
// 	age := make(map[string]int)
// 	age["mina"] = 28
// 	age["john"] = 32
// 	age["mike"] = 55
// 	fmt.Println("age of john:", age["john"])
// }

// --- initializing and accessing a map
// package main

// import "fmt"

// func main() {
// 	age := map[string]int{
// 		"mina": 28,
// 		"john": 32,
// 		"mike": 55,
// 	}

// 	minaAge, minaOk := age["mina"]
// 	jessyAge, jessyOk := age["jessy"]

// 	fmt.Println(minaAge, minaOk)
// 	fmt.Println(jessyAge, jessyOk)

// 	delete(age, "john")
// 	delete(age, "jessy")

// 	for key, value := range age {
// 		fmt.Println(key, "=>", value)
// 	}

// 	fmt.Println("len(age) =", len(age))
// }

// --- map with other data types
// package main

// import "fmt"

// func main() {
// 	age := map[bool]string{
// 		true: "YES",
// 		false: "NO",
// 	}

// 	for key, value := range age {
// 		fmt.Println(key, "=>", value)
// 	}
// }

// --- maps are referenced type
package main

import "fmt"

func main() {
	ages := make(map[string]int)

	age := map[string]int {
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	for key, value := range(age) {
		ages[key] = value
	}

	delete(ages, "john")

	fmt.Println("age", age)
	fmt.Println("ages", ages)
}