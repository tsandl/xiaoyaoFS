package storage

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	StoreDir string

	StoreAdminHost string
	StoreAdminPort int
	StoreApiHost string
	StoreApiPort int

	MasterHost string
	MasterPort int
}

// NewConfig new a config.
func NewConfig(conf string) (c *Config, err error) {
	var (
		file *os.File
		blob []byte
	)
	c = new(Config)
	if file, err = os.Open(conf); err != nil {
		return
	}
	if blob, err = ioutil.ReadAll(file); err != nil {
		return
	}
	err = toml.Unmarshal(blob, c)
	return
}