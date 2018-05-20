func reverseString(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		tmp := bytes[i]
		bytes[i] = bytes[j]
		bytes[j] = tmp
	}
	return string(bytes)
}
