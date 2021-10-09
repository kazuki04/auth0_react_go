package config

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Audience string
	Iss      string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("../config.ini")
	Config = ConfigList{
		Audience: cfg.Section("auth0").Key("audience").String(),
		Iss:      cfg.Section("auth0").Key("iss").String(),
	}
}
