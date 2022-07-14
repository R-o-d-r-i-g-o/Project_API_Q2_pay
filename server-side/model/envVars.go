package model

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_DATABASE"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Charset  string `env:"CHARSET"`
}

func (con *Config) GetConfig() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Evironment variables load fail (.env file)!")

	} else {
		env.Parse(con)
	}
}
