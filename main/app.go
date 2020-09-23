package main

import (
	"fmt"
	"github.com/fadhilahm/nummaniptwo/calc"
	calcNew "github.com/fadhilahm/nummaniptwo/v2/calc"
)

func main() {
	result := calc.Add(1, 2, 3)
	fmt.Println("calc.Add(1, 2, 3) => ", result)

	// v2 of nummanip
	err, newResult := calcNew.Add()

	if err != nil {
		fmt.Println("Error: => ", err)
	} else {
		fmt.Println("calcNew.Add(1, 2, 3, 4) => ", newResult)
	}
}