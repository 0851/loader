package utils

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"reflect"
	"strings"
)

func parseConfig(config *Config, i interface{}, paths ...string) error {
	if config.Debug {
		fmt.Println("start parse config")
	}
	if len(paths) <= 0 {
		return nil
	}
	datas, err := waits(paths, func(item string) ([]byte, error) {
		if strings.HasPrefix(item, "http") {
			return readAsHttp(item)
		}
		p, err := filepath.Abs(item)
		if err != nil {
			return nil, err
		}
		return readAsFile(p)
	})
	if err != nil {
		return err
	}
	for _, item := range datas {
		err := Parse(i, item, config.Debug)
		if err != nil {
			return err
		}
	}
	if config.Debug {
		fmt.Println("end parse config")
	}
	return nil
}
func ParseConfig(config *Config, i interface{}, paths ...string) error {
	value := reflect.Indirect(reflect.ValueOf(i))
	if !value.CanAddr() {
		return fmt.Errorf("data %v err", i)
	}
	err := parseConfig(config, i, paths...)
	if err != nil {
		return err
	}
	if config.Reload == true && config.Interval != 0 {
		Auto(config.Interval, func() (bool, error) {
			defer func() {
				if config.CallBack != nil {
					config.CallBack()
				}
			}()
			//set default value at parse before
			reflectPtr := reflect.New(reflect.ValueOf(i).Elem().Type())
			reflectPtr.Elem().Set(value)
			err := parseConfig(config, i, paths...)
			return true, err
		}, config.Debug)
	}
	return nil
}

func parseBase(i interface{}, content []byte, debug bool) error {
	err := json.Unmarshal(content, i)
	if err == nil {
		return nil
	}
	err = yaml.Unmarshal(content, i)
	if err == nil {
		return nil
	}
	err = toml.Unmarshal(content, i)
	if err == nil {
		return nil
	}
	if debug {
		Failure(err, "parse error")
	}
	return err
}

// 解析配置文件
func Parse(i interface{}, content []byte, debug bool) error {
	err := parseBase(i, content, debug)
	if err != nil {
		return err
	}
	_ = env.Parse(i)
	return nil
}
