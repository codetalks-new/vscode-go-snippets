import (
	"strconv"
	"strings"
)

func joinNums(arr []int, sep string) string {
	numCount := len(arr)
	if numCount == 0 {
		return ""
	}
	strs := make([]string, numCount)
	for i := 0; i < numCount; i++ {
		strs[i] = strconv.Itoa(arr[i])
	}
	return strings.Join(strs, sep)
}