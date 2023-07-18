package config

import (
	"goCode/consul"
	"goCode/db"
	"goCode/log"
)

type GlobalConfig struct {
	Service consul.Server `mapstructure:"server" json:"server" yaml:"server"`
	Zap     log.Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL   db.MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   db.Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
}
