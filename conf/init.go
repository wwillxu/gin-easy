package conf

import (
	"gineasy/models"
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppPort    string
	AppRunMode string

	JwtKey string
)

func init() {
	cfg, err := ini.Load("conf/example.ini")
	if err != nil {
		log.Fatal("Profile read failed")
	}

	AppPort = cfg.Section("app").Key("Port").String()
	AppRunMode = cfg.Section("app").Key("RunMode").String()

	JwtKey = cfg.Section("jwt").Key("Key").String()

	//
	models.DatabaseInit(cfg.Section("database").Key("ConnectUri").String())
}
