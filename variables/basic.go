package variables

import (
	"fmt"
	"strings"
)

// TODO: Declare and initialize two integers in two different ways
// TODO: Compute sum, difference, product, quotient, and remainder
// TODO: Swap first int's value and second int's value without a third variable
func IntFunc() {
	i := 10
	var j int = 2

	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	fmt.Println(i % j)

	i, j = j, i

	fmt.Println(i)
	fmt.Println(j)

}

// TODO: Declare and initialize two strings
// TODO: Create a third string, that is concataneted first two strings and "!"
// TODO: Make all characters in the third string upper case: use strings package
func StringFunc() {
	s1, s2 := "one ", "two"
	s3 := s1 + s2 + "!"
	fmt.Println("s3:", s3)

	fmt.Println(strings.ToUpper(s3))

}

// TODO: Declare and initialize booleans variables:
// 1. by comparing two integers using comparisons ==, !=, <, >, <=, >=
// 2. and using logical operators &&, ||, !
func BoolFunc() {

	b1 := 1 == 2
	b2 := 1 != 3
	b3 := 3 > 4
	b4 := 4 < 5
	b5 := 5 >= 1
	b6 := 5 <= 6
	b7 := b1 || b2
	b8 := b3 && b4
	b9 := !b6

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)
	fmt.Println("b4: ", b4)
	fmt.Println("b5: ", b5)
	fmt.Println("b6: ", b6)
	fmt.Println("b7: ", b7)
	fmt.Println("b8: ", b8)
	fmt.Println("b9: ", b9)

}

// TODO: Complete StructFunc() first
// TODO: Using the Building struct from before, create a pointer to a new Building
// TODO: Create a slice of pointers to buildings, which contain one pointer to a building
// TODO: Add the first pointer to the slice
// TODO: Change the value of first buildings street and print out the buildings
func PointerFunc() {
	build := &Building{
		Street:     "rigas",
		BuildingNr: 99,
	}
	slc := []*Building{
		&Building{
			Street:     "rigas",
			BuildingNr: 100,
		},
	}
	slc = append(slc, build)

	build.BuildingNr = 101

	for _, v := range slc {
		fmt.Println(v)
	}
}
