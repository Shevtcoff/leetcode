package leetcode

/*
Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/
func strStr(haystack string, needle string) int {
	for i := 0; i <= (len(haystack) - len(needle)); i++ {
		if needle == haystack[i:len(needle)+i] {
			return i
		}
	}
	return -1
}
