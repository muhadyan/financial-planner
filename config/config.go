package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
	"github.com/joho/godotenv"
)

type Configuration struct {
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbPort     string `env:"DB_PORT"`
	DbHost     string `env:"DB_HOST"`
}

func GetConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	configuration := Configuration{}

	err = gonfig.GetConf("", &configuration)
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	return configuration
}
