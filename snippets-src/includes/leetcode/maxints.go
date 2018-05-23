func maxInts(nums ...int) int {
	maxNum := -int(^uint(0)>>1) - 1
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}