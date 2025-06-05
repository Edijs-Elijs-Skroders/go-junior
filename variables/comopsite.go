package variables

import (
	"fmt"
)

// TODO: create an array of length 2 that holds integers
// TODO: iterate over each element of the array
// TODO: print out value and type of each element
func ArrayFunc() {

	arr := [2]int{2, 3} // todo add ints

	for i, v := range arr {
		fmt.Printf("elem %v, value %v, type %T \n", i, v, v)
	}

}

// TODO: create an empty slice that holds strings
// TODO: add string "hello" to it
// TODO: add string "world" to it twice
// tip: ... operator is used to unpack the elements of a second slice, if there is one :)
// TODO: remove last "world" form it
func SliceFunc() {
	slc2 := []string{
		"world",
	}

	slc := []string{
		"hello",
		"world",
	}

	slc = append(slc, slc2...)
	slc = slc[:2]
	fmt.Println(slc)
}

// TODO: create a map that has int keys and string values
// TODO: add a key and a value to it
// TODO: check if a key is empty
func MapFunc() {

	m := map[int]string{
		10: "ten",
	}
	m[7] = "zero"
	if _, ok := m[0]; !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println(" k")
	}
}

// TODO: Create a new struct Building with field Street as string and BuildingNr as int
// TODO: Create a new variable of type building
// TODO: Change the buildings address
// TODO: Print out the buildings full address
type StuctsName struct {
	Name   string
	Length int
}

type Building struct {
	Street     string
	BuildingNr int
}

func StructFunc() {
	building := Building{
		Street: "rigas",
		BuildingNr: 99,
	}

	building.BuildingNr = 60

	fmt.Printf("Building %v %v \n", building.Street, building.BuildingNr)
}
