package utils

import (
	"fmt"
	"sync"
	"time"
)

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/**
waits 并行执行
*/
func waits(paths []string, do func(item string) ([]byte, error)) ([][]byte, error) {
	type pathMutex struct {
		result [][]byte
		locker *sync.Mutex
		wait   *sync.WaitGroup
		errs   []error
	}
	lenPath := len(paths)
	chunks := pathMutex{
		locker: &sync.Mutex{},
		wait:   &sync.WaitGroup{},
	}
	chunks.wait.Add(lenPath)
	for _, item := range paths {
		go func(item string) {
			defer func() {
				chunks.locker.Unlock()
				chunks.wait.Done()
			}()
			chunks.locker.Lock()
			chunk, err := do(item)
			chunks.result = append(chunks.result, chunk)
			if err != nil {
				chunks.errs = append(chunks.errs, err)
			}

		}(item)
	}
	chunks.wait.Wait()
	if len(chunks.errs) > 0 {
		return nil, chunks.errs[0]
	}
	return chunks.result, nil
}

func run(i time.Duration, keep func() (bool, error), debug bool) {
	//wait := sync.WaitGroup{}
	//wait.Add(1)
	timer := time.AfterFunc(time.Second*i, func() {
		defer func() {
			//wait.Done()
		}()
		if debug {
			fmt.Println("start", time.Now())
		}
		isKeep, err := keep()
		if debug {
			fmt.Println("end", time.Now())
		}
		if err != nil {
			Failure(err, "time run error")
		}
		if isKeep {
			run(i, keep, debug)
		}
	})
	timer.Reset(time.Second * i)
	//wait.Wait()
}
func Auto(i time.Duration, keep func() (bool, error), debug bool) {
	run(i, keep, debug)
}
