package configs

import "os"

type Config struct {
	BasePath string
}

func GetConfig() *Config {
	wd, _ := os.Getwd()
	return &Config{
		BasePath: wd,
	}
}
