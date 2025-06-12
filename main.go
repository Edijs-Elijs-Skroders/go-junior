package main

import (
	"course/ds"
	"fmt"
)

func main() {
	words := []string{"hello", "hello", "world", "world", "world"}
	m := ds.WordCount(words)
	fmt.Println(ds.GetMostFrequent(m))

}
