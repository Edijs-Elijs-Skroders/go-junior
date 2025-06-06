package err

import (
	"fmt"
	"strconv"
)

// TODO: import strconv
// TODO: using for range loop, iterate over inputs and pass each input to function strconv.Atoi
// e.g. strconv.Atoi(input)
// TODO: check for the function returns. If it returns error, print it out, else print out the returned value
func ErrFunc() {
	inputs := []string{
		"4",
		"07",
		"abc",
		"-13",
		"1.5",
		"100",
		"foo123",
	}

	for _, i := range inputs{
		if v, err := strconv.Atoi(i); err != nil {
			fmt.Println(err)
		}else {
		fmt.Println(v)
		}
	}

}
