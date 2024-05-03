package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("a"))

}

/* Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.

Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6. */

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {

		if count > 0 && s[i] == ' ' {
			break
		} else if s[i] != ' ' {
			count++
		}
	}
	return count
}
