package palindromenumber

func isPalindrome(x int) bool {
	if 9 < x && x < 99 {
		return x/10 == x%10
	}
	return 0 <= x
}
