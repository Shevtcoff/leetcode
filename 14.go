package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Ответ", longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println("Ответ", longestCommonPrefix(strs))
	strs = []string{"ab", "a"}
	fmt.Println("Ответ", longestCommonPrefix(strs))
	strs = []string{"flower", "flower", "flower", "flower"}
	fmt.Println("Ответ", longestCommonPrefix(strs))
}

func longestCommonPrefix1(strs []string) string {
	var str string
	if len(strs) == 0 {
		return str
	} else if len(strs) == 1 {
		return strs[0]
	}
	first := strs[0]
	for k := 1; len(first) >= k; k++ {
		prefix := first[:k]
		for i := 1; i < len(strs); i++ {
			if k > len(strs[i]) || prefix != strs[i][:k] {
				return str
			}
		}
		str = prefix
	}
	return str
}
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for i := 1; len(strs) > i; i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
func longestCommonPrefix(strs []string) string {
	var str string
	if len(strs) == 0 {
		return str
	} else if len(strs) == 1 {
		return strs[0]
	}
	for k := 1; len(strs[0]) >= k; k++ {
		for i := 1; i < len(strs); i++ {
			if k > len(strs[i]) || strs[0][k-1] != strs[i][k-1] {
				return str
			}
		}
		str = strs[0][:k]
	}
	return str
}
