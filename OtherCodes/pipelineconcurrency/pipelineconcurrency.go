package pipelineconcurrency

import (
	"fmt"
)

func sliceToChannel(in []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, i := range in {
			out <- i
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func Main() {
	// input
	in := []int{2, 5, 6, 9, 1, 3}

	// stage 1
	dataChannel := sliceToChannel(in)

	// stage 2
	finalData := sq(dataChannel)

	//stage 3
	for n := range finalData {
		fmt.Println(n)
	}
}
