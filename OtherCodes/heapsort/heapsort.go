package heapsort

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func heap(arr []int, n int, pos int) {
	largest := pos

	left := largest*2 + 1
	right := largest*2 + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != pos {
		swap(&arr[largest], &arr[pos])
		heap(arr, n, largest)
	}

}

func heapsort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heap(arr, len(arr), i)

	}

	for i := len(arr) - 1; i > 0; i-- {
		swap(&arr[0], &arr[i])
		heap(arr, i, 0)
	}
}

func Main() {
	arr := []int{38, 27, 43, 10, -1, 12, -34, -2, -34, 23, 1, 233}
	heapsort(arr)
	fmt.Println(arr)
}
