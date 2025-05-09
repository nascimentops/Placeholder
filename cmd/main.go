package main

import (
	"fmt"

	multislices "github.com/nascimentops/multislices/internal/MultiSlices"
	stringconverter "github.com/nascimentops/multislices/internal/StringConverter"
)

func main() {

	y := multislices.MultiElements([]int{2, 3, 4, 5}, []int{3, 2, 1})
	x := stringconverter.ToBytes("banana")
	z := stringconverter.ToString(x)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(z)
}
