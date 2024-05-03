package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	mums2 := []int{2, 5, 6}
	merge(nums1, 3, mums2, 3)
	fmt.Println(nums1)
	nums1 = []int{1}
	mums2 = []int{}
	merge(nums1, 1, mums2, 0)
	fmt.Println(nums1)
	nums1 = []int{0}
	mums2 = []int{1}
	merge(nums1, 0, mums2, 1)
	fmt.Println(nums1)
}

/*

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

*/

func merge(nums1 []int, m int, nums2 []int, n int) {

	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[n+m-1] = nums1[m-1]
			m--
		} else {
			nums1[n+m-1] = nums2[n-1]
			n--
		}

	}

	fmt.Println("Результат", nums1)
}
