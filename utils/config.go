package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	AppName     string
	Port        string
	Debug       bool
	Limit       int
	PathLogging string
	DB          DatabaseCofig
}

type DatabaseCofig struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

func ReadConfiguration() (Configuration, error) {
	err := godotenv.Load()
	if err != nil {
		return Configuration{}, errors.New("Error loading .env file")
	}

	return Configuration{
		AppName:     os.Getenv("APP_NAME"),
		Port:        os.Getenv("PORT"),
		Debug:       StringToBool(os.Getenv("DEBUG")),
		Limit:       StringToInt(os.Getenv("LIMIT")),
		PathLogging: os.Getenv("PATH_LOGGING"),
		DB: DatabaseCofig{
			Name:     os.Getenv("DATABASE_NAME"),
			Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
		},
	}, nil

}
