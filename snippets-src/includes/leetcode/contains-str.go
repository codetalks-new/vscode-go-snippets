func findIndex(strs []string, str string) int {
	for index, curStr := range strs {
		if curStr == str {
			return index
		}
	}
	return -1
}

func containsStr(strs []string, str string) int {
	return findIndex(strs, str) != -1
}
