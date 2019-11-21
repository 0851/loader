package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

//
//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/BurntSushi/toml"
//	"gopkg.in/yaml.v3"
//	"io/ioutil"
//	"strings"
//)
//

func parseConfigSingle(i interface{}, paths ...string) error {
	datas, err := waits(paths, func(item string) ([]byte, error) {
		return nil, nil
	})
	if err != nil {
		return err
	}
	fmt.Println(datas)
	for _, item := range datas {
		Parse(i, item)
	}
	return nil
}
func ParseConfig(config *Config, i interface{}, paths ...string) error {
	defer func() {
		config.CallBack()
	}()
	err := parseConfigSingle(i, paths...)
	if config.Reload {
		Auto(config.Interval, func() (bool, error) {
			defer func() {
				config.CallBack()
			}()
			err := parseConfigSingle(i, paths...)
			return true, err
		}, config.Debug)
	}
	return err
}

// 解析配置文件
func Parse(i interface{}, content []byte) error {
	var err error = nil
	jsonErr := json.Unmarshal(content, &i)
	if jsonErr != nil {
		err = fmt.Errorf("json error :%s", jsonErr.Error())
	}
	yamlErr := yaml.Unmarshal(content, &i)
	if yamlErr != nil {
		err = fmt.Errorf("yaml error :%s", yamlErr.Error())
	}
	return err
}
