package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
}

/*
Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
*/
func mySqrt(x int) int {
	left, right := 0, x+1
	for left < right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
		fmt.Printf("mid = %d, left = %d, right = %d\n", mid, left, right)
	}
	return left - 1
}
