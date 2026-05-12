package main

import (
	"fmt"
	"go_essentials/any_processor"
	"go_essentials/bit_tricks"
	"go_essentials/complex"
	"go_essentials/constatns"
	"go_essentials/types"
)

func main() {
	types.TypesExample()
	bit_tricks.FastModPowerOfTwo(25, 3)
	bit_tricks.FastModPowerOfTwo(100, 4)
	fmt.Println(constatns.FormatBytes(1125899906842624))
	fmt.Println(complex.PrettyComplex(0 + 1i))

	data := [...]any{
		100, 12, 0, 10, 1000, "hello",
		"123", true, false, 3.14, "456", nil, []int{1, 2, 3, 4},
	}

	for i, v := range data {
		result := any_processor.ProcessAny(v)
		fmt.Printf("[%d] %v (%T) -> %v (%T) \n", i, v, v, result, result)
	}

}
