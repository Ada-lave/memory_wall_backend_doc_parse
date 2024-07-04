package config

import (
	"os"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
	Production *bool `yaml:"production" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var confPath string = os.Getenv("CONFIG_PATH")

	if confPath == "" {
		panic("ENV: Config path is empty")
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		panic(err)
	}

	var cnf Config

	if err := cleanenv.ReadConfig(confPath, &cnf); err != nil {
		panic(err)
	}

	return &cnf
}