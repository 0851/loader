package utils

import (
	"fmt"
	"time"
)

type Config struct {
	Debug    bool
	Interval time.Duration
	Reload   bool
	CallBack func()
}

func (config *Config) TraceTime(msg string) func() {
	var start time.Time
	if config.Debug == true {
		start = time.Now()
		fmt.Printf("enter %s\n", msg)
	}
	return func() {
		if config.Debug == true {
			fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
		}
	}
}
