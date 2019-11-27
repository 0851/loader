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
Waits 并行执行
*/

func Waits(l int, do func(item int) (interface{}, error)) ([]interface{}, error) {
	var result []interface{}
	var errs []error
	var wg sync.WaitGroup
	ch := make(chan chan interface{}, l)
	wg.Add(l)
	for i := 0; i < l; i++ {
		c := make(chan interface{}, 1)
		ch <- c
		go func(i int) {
			defer func() {
				wg.Done()
				close(c)
			}()
			chunk, err := do(i)
			c <- chunk
			if err != nil {
				errs = append(errs, err)
			}
		}(i)
	}
	wg.Wait()
	close(ch)
	for i := 0; i < l; i++ {
		c := <-ch
		result = append(result, <-c)
	}
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return result, nil
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
