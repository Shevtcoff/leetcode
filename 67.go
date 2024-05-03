package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

/*
*
Example 1:

Input: a = "11", b = "1"
Output: "100"

Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/
func addBinary(a string, b string) string {
	var res string

	indexA := len(a) - 1
	indexB := len(b) - 1
	shift := 0

	for indexA >= 0 || indexB >= 0 {
		x, y := 0, 0
		if indexA >= 0 {
			y = int(a[indexA] - '0')
			indexA--
		}
		if indexB >= 0 {
			y = int(b[indexB] - '0')
			indexB--
		}

		sum := strconv.Itoa(shift ^ x ^ y)
		shift = (shift & x) | (x & y) | (shift & y)

		res = sum + res
	}

	if shift == 1 {
		res = "1" + res
	}

	return res
}
