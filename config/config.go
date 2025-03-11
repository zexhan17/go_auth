package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from the .env file
func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}
