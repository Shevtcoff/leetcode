package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("cb"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("bbbbbb"))
}

/*
Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"
*/
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := string(s[0])
	for i := 0; i < len(s); i++ {
		j, k := i-1, i+1
		for k < len(s) && s[i] == s[k] {
			k++
			if len(res) < len(s[i:k]) {
				res = s[i:k]
			}
		}
		for j >= 0 && k < len(s) {
			if s[j] == s[k] {
				k++
				if len(res) < len(s[j:k]) {
					res = s[j:k]
				}
				j--
			} else {
				break
			}
		}
	}

	return res
}
