package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("("))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid(")(){}"))

}

/*
Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false
*/
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	stack := make([]rune, 0, len(s))
	hashBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		_, ok := hashBrackets[v]
		if !ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != hashBrackets[v] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
