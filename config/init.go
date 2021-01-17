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
	App app
	DB  db
	JWT jwt
	Dev bool
}

type app struct {
	Port   string
	Prefix string
}

type db struct {
	URI  string
	Name string
}

type jwt struct {
	Secret string
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
	if Conf.Dev {
		log.Println("[INFO] app running at dev mode")
	}
}
