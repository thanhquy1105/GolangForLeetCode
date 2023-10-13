package quicksort

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func quicksort(arr []int, from int, highest int) {
	if from >= highest {
		return
	}

	pos := from
	for i := from; i < highest; i++ {
		if arr[highest] < arr[i] {
			swap(&arr[pos], &arr[i])
			pos++
		}
	}
	swap(&arr[pos], &arr[highest])

	quicksort(arr, from, pos-1)
	quicksort(arr, pos+1, highest)

}

func Main() {
	arr := []int{38, 27, 43, 10, -1, 12, -34, -2, -34, 23, 1, 233}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
