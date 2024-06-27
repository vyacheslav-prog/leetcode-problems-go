package lettercombinationsofaphonenumber

var digitToLettersMap = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var combinations, result []string
	for _, digit := range digits {
		letters := digitToLettersMap[digit]
		combinations = nil
		for _, combination := range result {
			for _, letter := range letters {
				combinations = append(combinations, combination+letter)
			}
		}
		if nil == combinations {
			combinations = letters
		}
		result = combinations
	}
	return result
}
