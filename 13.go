package leetcode

/* func main() {
	fmt.Println("13. Roman to Integer")

	strNumber := "III"
	//	fmt.Printf("Example 1: %s = %d\n", strNumber, romanToInt(strNumber))
	print("Example 1: " + strNumber + " = ")
	println(romanToInt(strNumber))

	strNumber = "LVIII"
	fmt.Printf("Example 2: %s = %d\n", strNumber, romanToInt(strNumber))

	strNumber = "MCMXCIV"
	fmt.Printf("Example 3: %s = %d\n", strNumber, romanToInt(strNumber))
}
*/
/*
Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	mapSymbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	prevSymbol := rune(s[len(s)-1])
	arabNumber := mapSymbols[prevSymbol]
	for i := len(s) - 2; i > -1; i-- {
		char := rune(s[i])

		if mapSymbols[prevSymbol] <= mapSymbols[char] {
			arabNumber += mapSymbols[char]
		} else {
			arabNumber -= mapSymbols[char]
		}
		prevSymbol = char
	}

	return arabNumber
}
