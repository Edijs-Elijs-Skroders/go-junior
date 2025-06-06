package conditionals

import "fmt"

// TODO: create an int
// TODO: if this int is larger than 10, print out "larger"
// TODO: if its 10 or smaller, but larger or equal to 5, print "smaller, but larger or equal to 5"
// TODO: if its smaller than 5: print "smaller than 5"
func IfFunction() {
	i := 8
	if i > 10 {
		fmt.Println("larger")
	} else if i <= 10 || i >= 5 {
		fmt.Println("smaller, but larger or equal to 5")
	} else {
		fmt.Println("smaller than 5")
	}
}
