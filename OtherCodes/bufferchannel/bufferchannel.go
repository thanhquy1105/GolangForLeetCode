package bufferchannel

import (
	"fmt"
	"sync"
)

const (
	numberOfURLs = 100
	workers      = 5
)

var wg sync.WaitGroup

func crawlURLs(queue <-chan int, n int) {
	defer wg.Done()
	for i := range queue {
		// crawling
		fmt.Printf("worker %d is crawling URL %d\n", n, i)
	}

	fmt.Printf("Worker %d done and exit \n", n)
}

func startQueue() <-chan int {
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		close(queue)
	}()
	return queue
}

func Main() {
	// queue := startQueue()
	queue := make(chan int, 10)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go crawlURLs(queue, i)
	}

	for i := 0; i < numberOfURLs; i++ {
		queue <- i
		fmt.Printf("URL %d has been enqueued\n", i)
	}
	close(queue)

	wg.Wait()
}
