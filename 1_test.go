package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	target := 9
	nums := []int{2, 7, 11, 15}
	etalon := []int{0, 1}
	res := twoSum(nums, target)
	if !reflect.DeepEqual(res, etalon) {
		t.Errorf("Input: nums = %v, target = %d Output:  %v, res = %v", nums, target, etalon, res)
	}

	target = 6
	nums = []int{3, 2, 4}
	etalon = []int{1, 2}
	res = twoSum(nums, target)
	if !reflect.DeepEqual(res, etalon) {
		t.Errorf("Input: nums = %v, target = %d Output:  %v, res = %v", nums, target, etalon, res)
	}

	target = 6
	nums = []int{3, 3}
	etalon = []int{0, 1}
	res = twoSum(nums, target)
	if !reflect.DeepEqual(res, etalon) {
		t.Errorf("Input: nums = %v, target = %d Output:  %v, res = %v", nums, target, etalon, res)
	}
}
