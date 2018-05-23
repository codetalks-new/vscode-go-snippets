func minInts(nums ...int) int {
	minNum := int(^uint(0) >> 1)
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}