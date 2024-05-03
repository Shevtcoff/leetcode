package leetcode

/* func main() {
	fmt.Println("Example 1:")
	x := 121
	fmt.Println(isPalindrome(x))

	fmt.Println("Example 2:")
	x = -121
	fmt.Println(isPalindrome(x))

	fmt.Println("Example 3:")
	x = 10
	fmt.Println(isPalindrome(x))

	fmt.Println("Example 4:")
	x = 1234567654321
	fmt.Println(isPalindrome(x))

} */

/* Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome. */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	source := x
	var reversed int

	for x > 0 {
		reversed = (reversed * 10) + (x % 10)

		x = x / 10

		//fmt.Printf("r: %d, x: %d\n", reversed, x)
	}

	return reversed == source
}
