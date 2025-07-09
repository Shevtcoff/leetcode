package leetcode

/* func main() {

	target := 9
	nums := []int{2, 7, 11, 15}
	fmt.Println("Example 1:")
	fmt.Println(twoSum(nums, target))

	target = 6
	nums = []int{3, 2, 4}
	fmt.Println("Example 2:")
	fmt.Println(twoSum(nums, target))

	target = 6
	nums = []int{3, 3}
	fmt.Println("Example 3:")
	fmt.Println(twoSum(nums, target))

} */

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if in, ok := hashMap[target-nums[i]]; ok {
			res = append(res, in, i)
			break
		} else {
			hashMap[nums[i]] = i
		}
	}
	return res
}


func twoSum2(nums []int, target int) []int {
    if len(nums) < 2 {
        return []int{}
    }
    res := make([]int, 0, 2)

    for i:= 0; i < len(nums)-1; i++ {
        for k:=i+1; k<len(nums); k++{
            if nums[i] + nums[k] == target {
                return []int{i,k}
            }
        }
    }
    return res
}
