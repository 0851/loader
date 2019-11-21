package loader

import (
	"github.com/0851/loader/utils"
	"os"
)

type Config = utils.Config

type A struct {
	Config
}
type Loader struct {
	Config *Config
}

func New(config *Config) *Loader {
	if config == nil {
		config = &Config{}
	}
	if config.Interval == 0 {
		config.Interval = 10
	}
	if os.Getenv("LOADER_DEBUG") != "" {
		config.Debug = true
	}
	return &Loader{Config: config}
}

func Load(config *Config, i interface{}, paths ...string) (*Loader, error) {
	return New(config).Load(i, paths...)
}

func (loader *Loader) Load(i interface{}, paths ...string) (*Loader, error) {
	defer loader.Config.TraceTime("load")()
	utils.ParseConfig(loader.Config, i, paths...)
	return loader, nil
}
