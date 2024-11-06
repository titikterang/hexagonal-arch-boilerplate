package config

type Config struct {
	App   AppConfig   `json:"app"`
	Redis RedisConfig `json:"redis"`
}

type AppConfig struct {
	Port  int    `json:"port"`
	Label string `json:"label"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
