package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
}

/*
Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].

Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
*/
func plusOne(digits []int) []int {
	var shift int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + shift
		shift = 0

		if digits[i] >= 10 {
			println(digits[i])
			digits[i] = digits[i] % 10
			println("sum", digits[i])
			shift = 1
		} else {
			digits[i] = digits[i]
		}
	}
	println("shift", shift)
	if shift == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
