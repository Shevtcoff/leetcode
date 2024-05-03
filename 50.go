package main

import (
	"fmt"
)

/* Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000

Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100

Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25


    -100.0 < x < 100.0
    -231 <= n <= 231-1
    n is an integer.
    Either x is not zero or n > 0.
    -104 <= xn <= 104


*/

func main() {
	x, n := 2.0, 10
	fmt.Println(myPow(x, n))
	x, n = 2.1, 3
	fmt.Println(myPow(x, n))
	x, n = 2.0, -2
	fmt.Println(myPow(x, n))
	x, n = -1, 1147483648
	fmt.Println(myPow(x, n))
	fmt.Println(4 >> 1)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 == 1 {
		return x * myPow(x, n-1)
	}

	return myPow(x*x, n/2)
}
