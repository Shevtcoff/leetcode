package main

import "fmt"

func main() {

	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

/*
Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []


*/

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sumCandidates(candidates, target, []int{}, &res)

	return res
}

func sumCandidates(candidates []int, target int, current []int, res *[][]int) {
	for i, v := range candidates {
		if target-v <= 0 {
			if target-v == 0 {
				match := make([]int, len(current)+1)
				copy(match, current)
				match[len(current)] = v
				*res = append(*res, match)
			}

			break
		}
		fmt.Println(i, candidates[i:], target-v, append(current, v))
		sumCandidates(candidates[i:], target-v, append(current, v), res)
	}
}
