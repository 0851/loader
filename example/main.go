package main

import (
	"fmt"
	"github.com/0851/loader"
)

func main() {
	config := &loader.Config{
		Debug:  true,
		Reload: true,
		CallBack: func() {
			fmt.Println("这是一个测试")
		},
	}
	data := make(map[string]string)
	loader.Load(config, &data, "sdd", "sddd")
}
