package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(wg *sync.WaitGroup, jobs <-chan Task, results chan<- error) {
	defer wg.Done()
	for j := range jobs {
		err := j()
		results <- err
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// Solution based on https://gobyexample.com/worker-pools
func Run(tasks []Task, n, m int) error {

	var errCount int32
	var hitErrRate bool

	jobs := make(chan Task, len(tasks))
	results := make(chan error, len(tasks))

	wg := sync.WaitGroup{}

	// run workers
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}
	// send jobs
	var err error
	for j := 0; j < len(tasks); j++ {
		// https://gobyexample.com/non-blocking-channel-operations
		err = nil
		select {
		case err = <-results:
		default:
		}

		if err != nil {
			atomic.AddInt32(&errCount, 1)
			if atomic.LoadInt32(&errCount) >= int32(m) {
				hitErrRate = true
				break
			}
		}

		jobs <- tasks[j]
	}

	close(jobs)

	wg.Wait()

	if hitErrRate {
		return ErrErrorsLimitExceeded
	}

	return nil
}
