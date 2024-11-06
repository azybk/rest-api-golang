package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error When Load .env file: ", err.Error())
	}

	ExpInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
			Asset: os.Getenv("SERVER_ASSET_PATH"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TIMEZONE"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: ExpInt,
		},
		Storage: Storage{
			BasePath: os.Getenv("STORAGE_PATH"),
		},
	}
}
