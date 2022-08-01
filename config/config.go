package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit((1))
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitbank").Key("api_key").String(),
		ApiSecret: cfg.Section("bitbank").Key("api_secret").String(),
		LogFile:   cfg.Section("goTrading").Key("log_file").String(),
	}
}
