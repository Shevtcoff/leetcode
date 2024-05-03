package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("with words 4193"))
}

/**
Example 1:
Input: s = "42"
Output: 42

Example 2:
Input: s = "   -42"
Output: -42

Example 3:
Input: s = "4193 with words"
Output: 4193

Example 4:
Input: s = "with words 4193"
Output: 0
*/

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	resInt := 0
	mod := 1
	sign := false
	for _, char := range s {
		if !sign {
			if char == ' ' {
				continue
			} else if char == '+' {
				mod = 1
				sign = true
			} else if char == '-' {
				mod = -1
				sign = true
			} else if char >= '0' && char <= '9' {
				resInt = resInt*10 + int(char-'0')
				sign = true
			} else {
				return 0
			}
		} else if char >= '0' && char <= '9' {
			resInt = resInt*10 + int(char-'0')
			if resInt*mod <= math.MinInt32 {
				return math.MinInt32
			} else if resInt*mod >= math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	return resInt * mod
}
