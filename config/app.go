package config

type app struct {
	Addr string `yaml:"addr"`
	Env string `yaml:"env"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	App app `yaml:"app"`
}