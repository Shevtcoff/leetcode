package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for k := i + 1; k < len(nums); k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}
			l, r := k+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[k] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[k], nums[l], nums[r]})
					l++
					for ; l < len(nums)-1 && nums[l] == nums[l-1]; l++ {
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}

		}
	}
	return res
}
