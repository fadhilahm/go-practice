// --- what is string?
package main

import "fmt"

func main() {
	// rune
	// s := "Hellõ World"
	// r := []rune(s)

	// fmt.Println("len(r) =", len(r))

	// for i := 0; i < len(r); i++ {
	// 	fmt.Print(r[i], " ")
	// }

	// fmt.Println("")

	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%c ", r[i])
	// }

	// fmt.Println("")

	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%v ", r[i])
	// }

	// fmt.Println("")

	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%x ", r[i])
	// }

	// fmt.Println("")

	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%T ", r[i])
	// }

	// fmt.Println("")

	// using for loop in a string
	s := "Hellõ World"

	for index, char := range s {
		fmt.Printf("character at index %d is %c\n", index, char)
	}

	// what is rune
	r := 'õ'
	fmt.Printf("%x ", r)
	fmt.Printf("%v ", r)
	fmt.Printf("%T ", r)
	fmt.Println()

	// strings are immutable
	var1 := []uint8{72, 101, 108, 108, 111} // [72 101 108 108 111]
	var2 := string(var1) // Hello
	println(var1, var2)

	// string literals using backticks
	literal := `Hello \n
			My Big Blue
		"World"!`
	fmt.Println(literal)

	// character comparison
	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
	fmt.Printf("value of character b is %v of type %T\n", 'b', 'b')
	fmt.Println("hence 'b' > 'a' is", 'b' > 'a')

	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
	fmt.Printf("value of character A is %v of type %T\n", 'A', 'A')
	fmt.Println("hence 'A' > 'a' is", 'A' > 'a')

	fmt.Printf("\nvalue of character ℻ is %v of type %T\n", '℻', '℻')
	fmt.Printf("value of character ™ is %v of type %T\n", '™', '™')
	fmt.Println("hence '℻' > '™' is", '℻' > '™')

	// loop between two character value
	for i := 'a'; i < 'g'; i++ {
		fmt.Printf("character = '%c' with decimal value %v\n", i, i)
	}
}