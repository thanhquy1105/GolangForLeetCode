package countandsay

import (
	"fmt"
	"strconv"
)

func substrAndCout(s string) string {
	c := 0
	res := ""
	var temp string
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			temp = s[c:i]
			res += strconv.Itoa(len(temp)) + s[i-1:i]
			c = i
		}
	}
	temp = s[c:]
	res += strconv.Itoa(len(temp)) + s[len(s)-1:]
	return res
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return substrAndCout(countAndSay(n - 1))
}

func Main() {
	fmt.Println(countAndSay(5))
}
