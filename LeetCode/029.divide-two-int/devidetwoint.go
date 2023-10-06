package dividetwoint

import (
	"fmt"
	"math"
)

func abs(m int) (int, bool) {
	if m < 0 {
		return -m, false
	}
	return m, true
}
func divide(dividend int, divisor int) int {
	// single special case that would cause overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0
	divid, signDivid := abs(dividend)
	divis, signDivis := abs(divisor)
	for divid >= divis {
		sub := divis
		add := 1
		for divid >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		divid -= sub
		quotient += add
	}

	if signDivid != signDivis {
		return -quotient
	}

	return quotient
}

func Main() {
	fmt.Println(divide(10000000, 2))
}
