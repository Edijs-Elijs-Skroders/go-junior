package ds

import "fmt"

//TODO: Remove all the even numbers from the slice
//TODO: return a slice containing no even integers
func FilterOutEven(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if v % 2 != 0 {
			res = append(res, v)
		}
	}
	return res
}


//TODO: Remove all the duplicates from the slice
//TODO: return a slice with no duplicate integers
func RemoveDuplicates(nums []int) []int {
	check := make(map[int]bool, len(nums))
	res := []int{}

	for _, v := range nums {
		fmt.Println(check[v])
		if !check[v]{
			res = append(res, v)
			check[v] = true
		}
	}

	return res

}
