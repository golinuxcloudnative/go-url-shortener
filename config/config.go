package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Redis RedisCfg
}

type RedisCfg struct {
	Host string `mapstructure:"REDIS_HOST"`
	Port string `mapstructure:"REDIS_PORT"`
}

func NewConfig() (config *Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file does not exist:", err)
		}
		return
	}

	redisCfg := &RedisCfg{}

	err = viper.Unmarshal(&redisCfg)
	if err != nil {
		log.Fatalf("unable to unmarshal config %v", err)
		return
	}

	config = &Config{
		Redis: *redisCfg,
	}

	return
}
