package atoi

import (
	"fmt"
	"math"
)

func sign(s string) (string, int) {
	sign := 1
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	} else if s[0] == '+' {
		s = s[1:]
	}
	return s, sign
}

func trim(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return s[:i]
		}
	}
	return s
}

func overflow(i int) (int, bool) {
	if i > math.MaxInt32 {
		return math.MaxInt32, true
	}
	if i < math.MinInt32 {
		return math.MinInt32, true
	}
	return i, false
}
func myAtoi(s string) int {
	var n int
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	s, sign := sign(s)

	s = trim(s)

	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}

	if len(s) > 12 {
		if sign == 1 {
			return math.MaxInt32
		}
		return math.MinInt32
	}

	base := 1 * sign

	overfl := false
	for i := len(s) - 1; i >= 0; i-- {
		n, overfl = overflow(n + (int(s[i])-48)*base)
		if overfl {
			return n
		}
		base = base * 10
	}

	return n
}

func Main() {
	s := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"
	fmt.Printf("%d", myAtoi(s))
}
