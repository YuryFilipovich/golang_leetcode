var obj = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sArray := []rune(s)
	var number int
	for i := 0; i < len(sArray); i++ {
		currentSymbolVal := obj[sArray[i]]
		if i+1 < len(sArray) && currentSymbolVal < obj[sArray[i+1]] {
			number -= currentSymbolVal
		} else {
			number += currentSymbolVal
		}
	}
	return number
}

var obj = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sArray := []rune(s)

	var number int

	for i := range sArray {
		currentSymbolVal := obj[sArray[i]]
		if i+1 < len(sArray) && currentSymbolVal < obj[sArray[i+1]] {
			number -= currentSymbolVal
		} else {
			number += currentSymbolVal
		}
	}
	return number
}