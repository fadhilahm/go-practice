// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// define empty slice
// 	var s []int

// 	fmt.Println("s == nil", s == nil)

// 	// create an array of int
// 	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} 

// 	// create new slice
// 	s = a[2:4]

// 	// a[2] = 33
// 	// a[3] = 44

// 	// fmt.Println("length of a =", len(a))
// 	// fmt.Println("length of s =", len(s))
// 	// fmt.Println("capacity of s =", cap(s))

// 	// fmt.Println("address of a[2]", &a[2])
// 	// fmt.Println("address of s[0]", &s[0])
// 	// fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])

// 	// fmt.Println("before -> a[2] =", a[2])
// 	// s[0] = 33
// 	// fmt.Println("after -> a[2] =", a[2])

// 	newS := append(s, 55, 56)
// 	fmt.Printf("s=%v, newS=%v\n", s, newS)
// 	fmt.Printf("len=%d, cap=%d\n", len(newS), cap(newS))
// 	fmt.Printf("a=%v", a)
// } 

// ---

package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Printf("before -> s=%v\n", s)
	fmt.Printf("before -> a=%v\n", a)
	fmt.Printf("before -> len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println("%a[2] == &s[0] is", &a[0] == &s[0])

	s = append(s, 50, 60, 70, 80, 90, 100, 110)
	fmt.Printf("after -> s=%v\n", s)
	fmt.Printf("after -> a=%v\n", a)
	fmt.Printf("after -> len=%d, cap=%d\n", len(s), cap(s))
	fmt.Println("%a[2] == &s[0] is", &a[0] == &s[0])
}