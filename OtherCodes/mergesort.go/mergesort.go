package mergesort

import "fmt"

func merge(arr []int, from int, to int) {
	if from == to {
		return
	}

	mid := (to + from) / 2

	merge(arr, from, mid)
	merge(arr, mid+1, to)

	x := from
	y := mid + 1
	res := []int{}

	for x <= mid && y <= to {
		if arr[x] < arr[y] {
			res = append(res, arr[x])
			x++
		} else {
			res = append(res, arr[y])
			y++
		}
	}
	for x <= mid {
		res = append(res, arr[x])
		x++
	}
	for y <= to {
		res = append(res, arr[y])
		y++
	}

	for i := from; i <= to; i++ {
		arr[i] = res[i-from]
	}

}

func Main() {
	arr := []int{38, 27, 43, 10, -1, 12, -34, -2, -34, 23, 1, 233}
	merge(arr, 0, len(arr)-1)
	fmt.Print(arr)

}
