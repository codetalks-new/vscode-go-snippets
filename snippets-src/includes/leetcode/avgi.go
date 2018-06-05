func avg(nums []int) float64 {
	s := 0
	for _, num := range nums {
		s += num
	}
	return float64(s) / float64(len(nums))
}