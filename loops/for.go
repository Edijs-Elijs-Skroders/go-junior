package loops

import "fmt"

// TODO: create a for‐loop that declares and initializes an int i to 0,
// runs 10 iterations, and increases i by 3 at the end of each iteration
// TODO: print out each type and value of i
// TODO: if i is even, skip the printing out
// TODO: if i is larger than 15, stop the loop
func ForBasicLoop() {
 for i, count := 0,0; count < 10 ; i, count = i +3, count +1{
	if i > 15 {
		break
	}else if i % 2 == 0{
		continue
	}
	fmt.Printf("type %T, value %v \n", i,i)
 }
}

// TODO: Create a map with 4 key/value pairs
// TODO: Print out each key and value using for-range loop
func ForRangeLoop() {
	m := map[string]int{
		"one": 1,
		"two":2,
	}

	for k, v := range m {
		fmt.Printf("key %v, value %v \n", k, v )
	}
}

// TODO: Create a nested for-loop
// TODO: while in inner loop, break out of the outer loop
func NestedLoop() {

	OuterLoop:
	for i := 0; i < 10; i ++ {
		for j:= 0; j < 10; j++ {
			break OuterLoop
		}
	}
}
