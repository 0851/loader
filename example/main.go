package main

import (
	"fmt"
	"github.com/0851/loader"
	"time"
)

type A struct {
	Name string
	Ass  string
	Song interface{}
	Sdd  string
}

func main() {
	data := A{}
	config := &loader.Config{
		Debug:  true,
		Reload: true,
		CallBack: func() {
			fmt.Println(data)
			fmt.Println("这是一个测试")
		},
	}
	_, err := loader.Load(config, &data, "example/test.yaml", "example/test.json", "example/test.toml")
	fmt.Println(data, err)
	for {
		select {
		case <-time.After(10 * time.Minute):
			//todo
		}
	}
}
