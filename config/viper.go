package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Enum with dev,prod,stag and test
type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
	Stag Environment = "stag"
	Test Environment = "test"
)

func InitConfig(env Environment) Config {
	var config Config
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(string(env))
	v.AddConfigPath("./config/")
	v.AddConfigPath("config")

	// Handle errors for loading the config file
	if err := v.ReadInConfig(); err != nil {

		env := viper.GetString("app.env")

		fmt.Println(env)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("configuration file not found")
		} else {
			log.Fatal("error on parsing configuration file")
		}
	}

	// Unmarshal the config file in the struct
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal("error on parsing configuration file to struct")
	}

	// Watch the config file for changes
	v.WatchConfig()

	return config
}
