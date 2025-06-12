package ds

type Person struct {
    Name string
    Age  int
    City string
}


//TODO: find all the persons who are older than the minAge
//TODO: return them in a slice
func FilterPersonsByAge(people []Person, minAge int) []Person {
	results := []Person{}
	for _, v := range people{
		if v.Age >= minAge{
			results = append(results, v)
		}
	}
	return results
}