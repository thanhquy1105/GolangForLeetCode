package fractiontodecimal

import (
	"fmt"
	"strconv"
)

type node struct {
	value     int
	remainder int
	next      *node
}

func abs(i int) (int, bool) {
	if i < 0 {
		return -i, true
	}
	return i, false
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	absN, signN := abs(numerator)
	absD, signD := abs(denominator)

	s := ""
	if signD != signN {
		s += "-"
	}

	s += strconv.Itoa(absN / absD)

	remainder := absN % absD

	if remainder == 0 {
		return s
	}
	s += "."

	remainder *= 10

	value := (remainder) / absD
	remainder = remainder % absD

	head := &node{
		value:     value,
		remainder: remainder,
		next:      nil,
	}

	cur := head
	var repeated *node
	repeated = nil
loop:
	for remainder != 0 {
		remainder *= 10
		value = remainder / absD
		remainder = remainder % absD

		temp := head
		for ; temp != nil; temp = temp.next {
			if temp.remainder == remainder && value == temp.value {
				repeated = temp
				break loop
			}
		}
		newNode := &node{
			value:     value,
			remainder: remainder,
			next:      nil,
		}

		cur.next = newNode
		cur = cur.next
	}

	for temp := head; temp != nil; temp = temp.next {
		if repeated != nil {
			if repeated == temp {
				s += "("
			}
			s += strconv.Itoa(temp.value)
		} else {
			s += strconv.Itoa(temp.value)
		}
	}
	if repeated != nil {
		s += ")"
	}

	return s
}

func Main() {
	fmt.Println(fractionToDecimal(-50, 8))

}
