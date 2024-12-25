package stringutils

func Reversed(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
func FindLongest(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	longest := slice[0]
	for _, s := range slice[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}
