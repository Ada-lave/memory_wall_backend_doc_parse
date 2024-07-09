package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address    string `yaml:"address" env-required:"true"`
	Production *bool  `yaml:"production" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var confPath string = "config/local.yaml"

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		panic(err)
	}

	var cnf Config

	if err := cleanenv.ReadConfig(confPath, &cnf); err != nil {
		panic(err)
	}

	return &cnf
}
