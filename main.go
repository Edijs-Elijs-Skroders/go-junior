package main

import (
	"course/methods/shape"
	"fmt"
)

func main() {
	c := &shape.Circle{Radius: 2.1}
	r := &shape.Rectangle{Width: 2, Height: 1.5}

	fmt.Println(c)
	fmt.Println(r)
	
	slc := []shape.Shape{c, r}
	shape.ResizeAll(slc, 2)

	fmt.Println(c)
	fmt.Println(r)

}
