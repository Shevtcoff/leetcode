package main

import "fmt"

func main() {
	ar := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(ar)
	sortColors(ar)
	fmt.Println(ar)
	ar = []int{2, 0, 1}
	fmt.Println(ar)
	sortColors(ar)
	fmt.Println(ar)
}

/* Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]
*/

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		i := i
		for k := i; k < len(nums); k++ {
			if nums[k] < nums[i] {
				nums[k], nums[i] = nums[i], nums[k]
			}
		}
	}
}
