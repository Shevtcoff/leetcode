package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1122332211))
}

func reverse(x int) int {
	reverseX := 0

	for x != 0 {
		if reverseX > math.MaxInt32/10 || reverseX < math.MinInt32/10 {
			return 0
		}
		reverseX = (x % 10) + reverseX*10
		x = x / 10
	}
	return reverseX

}
