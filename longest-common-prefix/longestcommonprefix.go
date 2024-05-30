package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	var prefix string
	var sliceIndex int
	for index, str := range strs {
		if 0 == index {
			prefix = str
		}
		sliceIndex = 0
		for sliceIndex != len(prefix) && sliceIndex != len(str) && str[sliceIndex] == prefix[sliceIndex] {
			sliceIndex += 1
		}
		prefix = str[:sliceIndex]
	}
	return prefix
}
