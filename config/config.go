package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbPort     string `env:"DB_PORT"`
	DbHost     string `env:"DB_HOST"`

	SendFromAddress string `env:"SEND_FROM_ADDRESS"`
	MailPassword    string `env:"MAIL_PASSWORD"`

	BaseUrl string `env:"BASE_URL"`
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
