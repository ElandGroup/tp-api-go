package config

import "github.com/jinzhu/configor"

var Config = struct {
	Sample struct {
		Conn string `json:"Conn"`
	} `json:"Sample"`
}{}

func InitConfig(cfg string) {
	// configor.Load(&Config, "config/config.json")
	configor.Load(&Config, cfg)
}
