package shellsort

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func inssort(arr []int, n int, incr int) {
	for i := incr; i < n; i += incr {
		for j := i; (j >= incr) && (arr[j] > arr[j-incr]); j -= incr {
			swap(&arr[j], &arr[j-incr])
			fmt.Println(arr)
		}
	}
}

func shellsort(arr []int, n int) {

	for i := n / 2; i > 2; i /= 2 {
		for j := 0; j < i; j++ {
			inssort(arr[j:], n-j, i)
		}
	}
	inssort(arr, n, 1)
}

func Main() {
	arr := []int{-2, 3, -6, 1, 7, 4, 5}
	fmt.Println(arr)
	shellsort(arr, len(arr))
	fmt.Println(arr)

}
