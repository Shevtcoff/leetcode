package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1}, 1))
}

/* Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func searchInsert(nums []int, target int) int {
	leftIndex, rightIndex := 0, len(nums)
	for leftIndex < rightIndex {
		midIndex := (leftIndex + rightIndex) / 2
		if nums[midIndex] == target {
			return midIndex
		} else if nums[midIndex] < target {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex
		}
	}
	return leftIndex
}
