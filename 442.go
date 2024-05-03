package leetcode

import (
	"math"
)

/*
Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]

Example 2:

Input: nums = [1,1,2]
Output: [1]

Example 3:

Input: nums = [1]
Output: []
*/
/* func main() {
	fmt.Println("Example 1:")
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))

	fmt.Println("Example 2:")
	nums = []int{1, 1, 2}
	fmt.Println(findDuplicates(nums))

	fmt.Println("Example 3:")
	nums = []int{1}
	fmt.Println(findDuplicates(nums))

	fmt.Println("Example 4:")
	nums = []int{3, 11, 8, 16, 4, 15, 4, 17, 14, 14, 6, 6, 2, 8, 3, 12, 15, 20, 20, 5}
	fmt.Println(findDuplicates(nums))

	fmt.Println("Example 5:")
	nums = []int{1, 1}
	fmt.Println(findDuplicates(nums))

} */

func findDuplicates(nums []int) []int {

	res := make([]int, 0)
	for i := range nums {
		key := int(math.Abs(float64(nums[i])) - 1)
		if nums[key] < 0 {
			res = append(res, key+1)
		}
		nums[key] = -nums[key]
	}

	return res
}
