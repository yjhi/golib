package jconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	_configFile string
	_configType string
	_fileData   []byte
}

func BuildConfig(f string, t string) (*Config, error) {
	cfg := &Config{
		_configFile: f,
		_configType: t,
	}

	var err error
	cfg._fileData, err = ioutil.ReadFile(cfg._configFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
