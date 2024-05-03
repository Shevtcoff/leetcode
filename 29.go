package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(1<<1, 1<<2, 1<<3)
	fmt.Println(divide(100, 3))
}

func divide(dividend int, divisor int) int {
	if math.MinInt32 == dividend && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 {
		dividend = -dividend
		sign = -1
	} else if divisor < 0 {
		divisor = -divisor
		sign = -1
	}
	res := 0
	fmt.Println(dividend / divisor)
	for dividend >= divisor {
		multiple := divisor
		for i := 0; dividend >= multiple; i++ {
			dividend -= multiple
			res += 1 << uint(i)
			fmt.Printf("i=%d res=%d multipe=%d dividend=%d\n", i, res, multiple, dividend)
			multiple <<= 1
		}
	}
	return res * sign
}

func divide2(dividend int, divisor int) int {
	if math.MinInt32 == dividend && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 {
		dividend = -dividend
		sign = -1
	} else if divisor < 0 {
		divisor = -divisor
		sign = -1
	}
	res := 0
	for dividend >= divisor {

		multiple := divisor
		origDividend := dividend
		for i := 0; origDividend >= multiple; i++ {
			dividend -= multiple
			multiple <<= 1
			res = 1 << uint(i)
			fmt.Printf("i=%d res=%d multipe=%d dividend=%d\n", i, res, multiple, dividend)
			if origDividend-multiple < divisor {

				break
			}
		}
	}
	return res * sign
}
