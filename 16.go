package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := nums[0] + nums[1] + nums[2]
	result := sum
	difference := int(math.Abs(float64(result - target)))

	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]

			if sum-target < 0 {
				j++
			} else {
				k--
			}

			if difference > int(math.Abs(float64(sum-target))) {
				difference = int(math.Abs(float64(sum - target)))
				result = sum
			}
		}
	}

	return result
}
