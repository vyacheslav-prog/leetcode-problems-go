package integertoroman

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
	for numOfMap, symbol := range numberToSymbolMap {
		if remainder := num - numOfMap; 0 < remainder {
			return symbol + intToRoman(remainder)
		}
	}
	return ""
}
