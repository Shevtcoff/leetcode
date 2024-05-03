package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 5))
	fmt.Println(convert("A", 1))
}

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);



Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

Example 3:

Input: s = "A", numRows = 1
Output: "A"

*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); {
			result.WriteByte(s[j])
			j += numRows*2 - 2
			if i == 0 || i == numRows-1 {
				continue
			}
			j -= i + i
			if len(s) <= j {
				break
			}
			result.WriteByte(s[j])
			j += i + i
		}
	}

	return result.String()
}
