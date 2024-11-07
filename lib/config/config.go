package config

import "time"

type Config struct {
	App   AppConfig   `json:"app" mapstructure:"app"`
	Http  HttpConfig  `json:"http" mapstructure:"http"`
	Redis RedisConfig `json:"redis" mapstructure:"redis"`
}

type AppConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Label   string `json:"label" mapstructure:"label"`
}

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	PoolSize int    `json:"poolsize" mapstructure:"poolsize"`
}
