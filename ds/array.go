package ds

type Stats struct {
	Sum     int
	Min     int
	Max     int
	Average float64
}

//TODO: calculate sum of array elements
//TODO: find the smallest integer in the array
//TODO: find the largest integer in the array
//TODO: calculate the average from the integers in array
//TODO: return Stats, that has all of the mentioned above
func CalculateStats(arr [10]int) Stats {
	s := Stats{
		Min: arr[0],
		Max: arr[0],
	}

	for _, v := range arr {
		s.Sum = s.Sum + v
		if s.Max < v {
			s.Max = v
		}
		if s.Min > v {
			s.Min = v
		}
	}
	s.Average = float64(s.Sum) / float64(len(arr))

	return s
}