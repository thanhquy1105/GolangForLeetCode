package reverseint

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	sign := 1
	if x < 0 {
		sign = -1
		s = s[1:]
	}
	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		x += sign * (int(s[i]) - 48) * int(math.Pow(10, float64(i)))

		if x > math.MaxInt32 || x < math.MinInt32 {
			return 0
		}
	}
	return x

}

func Main() {
	fmt.Println(reverse(563847412))
}
