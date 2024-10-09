package config

type app struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	App app `yaml:"app"`
}