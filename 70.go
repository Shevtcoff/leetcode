package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(6))
}

/*
Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


*/

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	curr, next := 2, 3
	for i := 2; i < n; i++ {
		curr, next = next, curr+next
	}

	return curr

}
