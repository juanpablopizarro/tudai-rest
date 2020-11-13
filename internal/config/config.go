package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig ...
type DbConfig struct {
	Type   string `yaml:"type"`
	Driver string `yaml:"driver"`
	Conn   string `yaml:"conn"`
}

// GreeterConfig ...
type GreeterConfig struct {
	Version string `yaml:"version"`
}

// Config ...
type Config struct {
	DB      DbConfig      `yaml:"db"`
	Version string        `yaml:"version"`
	Greeter GreeterConfig `yaml:"greeter"`
}

// LoadConfig ...
func LoadConfig(filename string) *Config {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}

	return c
}
