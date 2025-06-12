package ds


//TODO: create a map with string key and integer value
//TODO: Find out how many times this word repeats in the slice and map the word to its count in a map
//TODO: return this map
func WordCount(words []string) map[string]int {
	m := make(map[string]int, len(words))

	for _, v := range words {
		m[v] = m[v] + 1
	}

	return m
}


//TODO: pass the map you got from the function above to this function
//TODO: Find out which words/keys have the highest count and return the key in a slice
func GetMostFrequent(counts map[string]int) []string {
	max := 0
	
	results := []string{}

	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	for k,v := range counts {
		if v == max {
			results = append(results, k)
		}
	}

	return results
}