package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	/*fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("ua"))
	fmt.Println(lengthOfLongestSubstring("dvkds"))*/

}

func lengthOfLongestSubstring(s string) int {
	longest, start := 0, 0
	seen := map[rune]int{} // could use map[byte]int instead and use seen[s[i]] instead

	for i, ch := range s {
		if _, ok := seen[ch]; ok && seen[ch] >= start { // as it has been seen
			start = seen[ch] + 1 //+1 потому что индекс меньше длины
		}
		longest = max(i-start+1, longest) //+1 потому что индекс  меньше длины
		seen[ch] = i
	}

	return longest
}
