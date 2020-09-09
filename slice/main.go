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

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	s := a[2:4]
// 	fmt.Printf("before -> s=%v\n", s)
// 	fmt.Printf("before -> a=%v\n", a)
// 	fmt.Printf("before -> len=%d, cap=%d\n", len(s), cap(s))
// 	fmt.Println("%a[2] == &s[0] is", &a[2] == &s[0])

// 	s = append(s, 50, 60, 70, 80, 90, 100, 110)
// 	fmt.Printf("after -> s=%v\n", s)
// 	fmt.Printf("after -> a=%v\n", a)
// 	fmt.Printf("after -> len=%d, cap=%d\n", len(s), cap(s))
// 	fmt.Println("%a[2] == &s[0] is", &a[2] == &s[0])
// }

// ---

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6}
// 	s = append(s, 7, 8)

// 	fmt.Println("s=", s)
// 	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
// }

// --- copy

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var s1 []int
// 	s2 := []int{1, 2, 3}
// 	s3 := []int{4, 5, 6, 7}
// 	s4 := []int{1, 2, 3}

// 	n1 := copy(s1, s2)
// 	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
// 	fmt.Println("s1 == nil", s1 == nil)

// 	n2 := copy(s2, s3)
// 	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)

// 	n3 := copy(s3, s4)
// 	fmt.Printf("n3=%d, s3=%v, s4=%v", n3, s3, s4)
// }

// -- make

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s1 := make([]int, 2, 4)
// 	s2 := []int{1, 2, 3}

// 	fmt.Printf("before => s1=%v, s2=%v\n", s1, s2)
// 	fmt.Println("before => s1 == nil", s1 == nil)

// 	n1 := copy(s1, s2)
// 	fmt.Printf("after => n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
// 	fmt.Println("after => s1 == nil", s1 == nil)
// }

// -- unpack operator

// package main

// import "fmt"

// func main() {
// 	s1 := make([]int, 0, 10)
// 	fmt.Println("before, s1 =>", s1)

// 	s2 := []int{1,2,3}
// 	s1 = append(s1, s2...)
// 	fmt.Println("after, s1 =>", s1)

// }

// -- extract operator

// package main

// import "fmt"

// func main() {
// 	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	fmt.Println("s[:]", s[:])
// 	fmt.Println("s[2:]", s[2:])
// 	fmt.Println("s[:4]", s[:4])
// 	fmt.Println("s[2:4]", s[2:4])
// }

// -- pass by reference
// package main

// import "fmt"

// func makeSquares(array [10]int) {
// 	for index, elem := range array {
// 		array[index] = elem * elem
// 	}
// }

// func main() {
// 	s := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	makeSquares(s)

// 	fmt.Println(s)
// }

// -- delete slice elements

package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// delete element at index 2 ( == 2)
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}