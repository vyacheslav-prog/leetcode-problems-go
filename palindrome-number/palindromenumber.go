package palindromenumber

const shift = 10

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	decimals := 1
	for remainder := x / shift; remainder != 0; remainder /= shift {
		decimals *= shift
	}
	for remainder := x; decimals != 0; decimals /= shift * shift {
		if remainder/decimals != remainder%shift {
			return false
		}
		if decimals/(shift*shift) != 0 {
			remainder %= decimals
			remainder /= shift
		}
	}
	return true
}
