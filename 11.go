package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

/*
*
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49

Example 2:
Input: height = [1,1]
Output: 1
*/
func maxArea(height []int) int {
	i, k := 0, len(height)-1
	maxArea := 0
	for i != k {
		h := height[i]
		if height[k] < height[i] {
			h = height[k]
		}
		if maxArea < h*(k-i) {
			maxArea = h * (k - i)
		}
		if height[i] < height[k] {
			i++
		} else {
			k--
		}
	}

	return maxArea
}
