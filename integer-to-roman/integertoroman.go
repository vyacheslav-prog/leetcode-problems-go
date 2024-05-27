package integertoroman

var numberToSymbolKeys = []int{1000, 900, 100, 10, 9, 5, 4, 1, 0}

var numberToSymbolMap = map[int]string{
	1000: "M",
	900:  "CM",
	100:  "C",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
	0:    "",
}

func intToRoman(num int) string {
	if symbol, hasSymbol := numberToSymbolMap[num]; hasSymbol {
		return symbol
	}
	for _, numOfMap := range numberToSymbolKeys {
		symbol := numberToSymbolMap[numOfMap]
		if remainder := num - numOfMap; 0 < remainder {
			return symbol + intToRoman(remainder)
		}
	}
	return ""
}
