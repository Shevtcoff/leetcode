package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 5, 4, 6, 4, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		sort.Ints(nums)
		fmt.Println("reverse")
		return
	}

	j := len(nums) - 1
	for nums[j] <= nums[i] {
		j--
	}

	fmt.Println("nums, i and j: ", nums, i, j)
	nums[i], nums[j] = nums[j], nums[i]
	fmt.Println("nums, i and j: ", nums, i, j)
	reverse(nums[i+1:])

}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
