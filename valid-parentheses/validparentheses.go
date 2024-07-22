package validparentheses

var openToCloseBracket = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	var closeStack []byte
	for index := 0; index != len(s); index += 1 {
		closeBracket, isOpen := openToCloseBracket[s[index]]
		if isOpen {
			closeStack = append(closeStack, closeBracket)
		} else if 0 != len(closeStack) && s[index] == closeStack[len(closeStack)-1] {
			closeStack = closeStack[:len(closeStack)-1]
		} else {
			closeStack = append(closeStack, closeBracket)
		}
	}
	return 0 == len(closeStack)
}
