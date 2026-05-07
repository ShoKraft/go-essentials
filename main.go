package main

import (
	"fmt"
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

}
