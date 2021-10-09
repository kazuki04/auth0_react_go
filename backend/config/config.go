package config

import (
	"gopkg.in/ini.v1"

)

type ConfigList {
	Audience string
	Iss      string
}

var Config ConfigList

func init(){
	cfg, _ := ini.Load("../config.ini")
	Config = ConfigList{
		Audience: cfg.Section.Key("audience").String(),
		Iss: cfg.Section.Key("iss").String(),
	}
}
