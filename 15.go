package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}

func threeSum2(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int
	lenSlice := len(nums)
	if lenSlice < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < lenSlice-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for k := i + 1; k < lenSlice-1; k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}
			for j := k + 1; j < lenSlice; j++ {
				if j > k+1 && nums[j] == nums[j-1] {
					continue
				}
				if nums[i]+nums[k]+nums[j] == 0 {
					singleres := []int{nums[i], nums[k], nums[j]}
					res = append(res, singleres)

				}
			}
		}
	}

	return res
}
