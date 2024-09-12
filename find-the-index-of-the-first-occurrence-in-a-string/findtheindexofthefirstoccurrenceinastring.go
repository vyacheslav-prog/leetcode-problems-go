package findtheindexofthefirstoccurrenceinastring

const outOfBounds = -1

func strStr(haystack string, needle string) int {
	if 0 == len(haystack) || 0 == len(needle) || len(haystack) < len(needle) {
		return outOfBounds
	}
	firstIndex := outOfBounds
	for haystackCounter := 0; outOfBounds == firstIndex && haystackCounter != len(haystack)-len(needle)+1; haystackCounter += 1 {
		matchingIndex := haystackCounter
		for len(needle) != matchingIndex-haystackCounter && haystack[matchingIndex] == needle[matchingIndex-haystackCounter] {
			matchingIndex += 1
		}
		if matchingLength := matchingIndex - haystackCounter; 0 != matchingLength && len(needle) == matchingLength {
			firstIndex = haystackCounter
		}
	}
	return firstIndex
}
