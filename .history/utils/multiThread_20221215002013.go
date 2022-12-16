package utils

import "time"

func MultiThread() {
	multiThread()
}

func multiThread() {
	ch := make(chan time.Duration)

	go hugeIter(ch)
}

func hugeIter(ch chan time.Duration) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
	}
	elapse := time.Since(start)

	return elapse
}
