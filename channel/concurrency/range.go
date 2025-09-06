package main

func concurrentFib(n int) []int {
	// ?
	ch := make(chan int)
	fibSequence := make([]int, 0, n)
	go fibonacci(n, ch)
	for item := range ch {
		fibSequence = append(fibSequence, item)
	}
	return fibSequence
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
