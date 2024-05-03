package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	generatePairs("", 0, 0, n, &res)

	return res
}

func generatePairs(s string, left, right, n int, res *[]string) {

	if left+right == 2*n {
		*res = append(*res, s)
		return
	}

	if left < n {
		generatePairs(s+"(", left+1, right, n, res)
	}

	if right < left {
		generatePairs(s+")", left, right+1, n, res)
	}

}
