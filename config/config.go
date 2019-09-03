// Package config contains the configuration structure
package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

// Config holds all configuration values
type Config struct {
	Locale string
}

// NewConfig returns the config from the file
func NewConfig() (conf Config, err error) {
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return
	}

	_, err = toml.Decode(string(data), &conf)

	return
}
