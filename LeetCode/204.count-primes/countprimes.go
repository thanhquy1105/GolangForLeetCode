package countprimes

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	count := 1
	arr := make([]bool, n)

	for i := 3; i < n; i += 2 {
		if !arr[i] {
			count++
			for j := i * i; j < n; j += 2 * i {
				arr[j] = true
			}
		}
	}

	return count
}

func Main() {
	fmt.Println(countPrimes(16))
}
