package integertoroman

var numberToSymbolMap = map[int]string{
	0: "",
	1: "I",
	5: "V",
}

func intToRoman(num int) string {
	return numberToSymbolMap[num]
}
