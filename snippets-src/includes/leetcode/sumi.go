func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}