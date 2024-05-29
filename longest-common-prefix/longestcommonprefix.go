package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	var prefix string
	for index, str := range strs {
		if 0 == index || prefix == str {
			prefix = str
		} else {
			prefix = ""
		}
	}
	return prefix
}
