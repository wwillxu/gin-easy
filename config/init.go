package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Conf *Config
)

type Config struct {
	App app `yaml:"app"`
	DB  db  `yaml:"db"`
	JWT jwt `yaml:"jwt"`
}

type app struct {
	Port   string `yaml:"port"`
	Prefix string `yaml:"prefix"`
}

type db struct {
	URI  string `yaml:"uri"`
	Name string `yaml:"name"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}

func init() {
	configFile := "example.yml"
	data, err := ioutil.ReadFile(fmt.Sprintf("config/%s", configFile))

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatal("Error loading config.env file", err)
	}

	Conf = config
}
