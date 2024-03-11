package utils

import (
	"log"

	"github.com/spf13/viper"
)

type configStruct struct {
	DbUser string `mapstructure:"DBUSER"`
	DbPass string `mapstructure:"DBPASS"`
	DbName string `mapstructure:"DBNAME"`
	DbIP   string `mapstructure:"DBIP"`
	Port   string `mapstructure:"PORT"`
}

var config *configStruct

func InitConfig() configStruct {
	viper.SetConfigFile("local.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error while reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error while reading env file", err)
	}

	return *config
}

func GetConfig() configStruct {
	if config != nil {
		return *config
	}

	newConfig := InitConfig()
	return newConfig
}
